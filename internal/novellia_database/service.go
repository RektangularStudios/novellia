package novellia_database

import (
	"fmt"
	"context"
	"io/ioutil"
	"path/filepath"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4/pgxpool"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	"github.com/RektangularStudios/novellia/internal/constants"
	"github.com/RektangularStudios/novellia/internal/config"
	prometheus_monitoring "github.com/RektangularStudios/novellia/internal/monitoring"
)

const (
	queryProductID = "queryProductID"
	queryProduct = "queryProduct"
	queryCommission = "queryCommission"
	queryAttribution = "queryAttribution"
	queryRemoteResource = "queryRemoteResource"
)

type ServiceImpl struct {
	queriesPath string
	pool *pgxpool.Pool
	queries map[string]string
}

// creates a new ServiceImpl, connecting to Postgres
func New(ctx context.Context) (*ServiceImpl, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config")
	}

	// url like "postgresql://username:password@localhost:5432/database_name?search_path=novellia"
	schema := "novellia"
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s?search_path=%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.CardanoDatabase,
		schema,
	)
	
	pool, err := pgxpool.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Postgres: %v", err)
	}
	
	service := ServiceImpl {
		pool: pool,
		queriesPath: cfg.Postgres.QueriesPath,
	}
	err = service.loadQueries(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load queries")
	}

	return &service, nil
}

func (s *ServiceImpl) loadQueries(ctx context.Context) error {
	queryFiles := map[string]string {
		queryProductID: "query_product_id.sql",
		queryProduct: "query_product.sql",
		queryCommission: "query_commission.sql",
		queryAttribution: "query_attribution.sql",
		queryRemoteResource: "query_remote_resource.sql",
	}

	queries := make(map[string]string)
	for name, filename := range queryFiles {
		fmt.Printf("Loading SQL %s\n", filename)

		query, err := s.readQueryFile(filename)
		if err != nil {
			return err
		}

		queries[name] = query
	}
	s.queries = queries

	fmt.Printf("SQL has been loaded\n")
	return nil
}

// reads a text file using the queriesPath as the base path
func (s *ServiceImpl) readQueryFile(filename string) (string, error) {
	queryPath := filepath.Join(s.queriesPath, filename)

	bytes, err := ioutil.ReadFile(queryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read query file %s: %v", filename, err)
	}

	return string(bytes), nil
}

// queries product list matching market and organization filters
func (s *ServiceImpl) QueryProductIDs(ctx context.Context, organizationId string, marketId string) ([]nvla.ProductListElement, error) {
	rows, err := s.pool.Query(ctx, s.queries[queryProductID], organizationId, marketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []nvla.ProductListElement{}
	for rows.Next() {
		var p nvla.ProductListElement
		var modified pgtype.Timestamptz

		err = rows.Scan(
			&p.ProductId,
			&p.NativeTokenId,
			&modified,
		)
		if err != nil {
			return nil, fmt.Errorf("query product IDs failed: %v", err)
		}

		m := modified.Get()
		if m != nil {
			p.Modified = modified.Time.UTC().Format(constants.ISO8601DateFormat)
		}

		// add product element to slice
		products = append(products, p)
	}

	prometheus_monitoring.RecordNumberOfProductIDsListed(len(products))

	return products, nil
}

// queries product information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddProduct(ctx context.Context, productElements []nvla.ProductListElement) ([]nvla.Product, error) {
	productIDs := []string{}
	nativeTokenIDs := []string{}
	for _, product := range productElements {
		if product.ProductId != "" {
			productIDs = append(productIDs, product.ProductId)
		}
		if product.NativeTokenId != "" {
			nativeTokenIDs = append(nativeTokenIDs, product.NativeTokenId)
		}
	}

	rows, err := s.pool.Query(ctx, s.queries[queryProduct], productIDs, nativeTokenIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	products := []nvla.Product{}
	for rows.Next() {
		var p nvla.Product
		var t nvla.NovelliaStandardToken
		p.Product.NovelliaStandardToken = &t
		var dateListed pgtype.Timestamptz
		var dateAvailable pgtype.Timestamptz

		err = rows.Scan(
			// product
			&p.Product.ProductId, &p.Product.NovelliaStandardToken.Name,
			&p.Pricing.PriceCurrencyId, &p.Pricing.PriceUnitAmount, &p.Pricing.MaxOrderSize,
			&dateListed, &dateAvailable,
			// organization
			&p.Organization.OrganizationId, &p.Organization.Name, &p.Organization.Description,
			// market
			&p.Market.MarketId, &p.Market.Name, &p.Market.Description,
			// native token
			&t.NativeToken.PolicyId, &t.NativeToken.AssetId,
			// product detail
			&t.Copyright,
			&t.Publisher,
			&t.Version,
			&t.Id,
			&t.Tags,
			&t.Description.Short,
			&t.Description.Long,
			// if this is a token, these need to filled in by querying Cardano. this is beyond the scope of this service
			&p.Stock.Available,
			&p.Stock.TotalSupply,
		)
		if err != nil {
			return nil, fmt.Errorf("query and add product failed: %v", err)
		}

		// convert dates to strings
		p.Metadata.DateListed = dateListed.Time.UTC().Format(constants.ISO8601DateFormat)
		p.Metadata.DateAvailable = dateAvailable.Time.UTC().Format(constants.ISO8601DateFormat)

		// add product to slice
		products = append(products, p)
	}

	return products, nil
}

// queries commission information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddCommission(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.pool.Query(ctx, s.queries[queryCommission], productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c nvla.Commission
		var product_id string
		err = rows.Scan(
			&product_id,
			&c.Name,
			&c.Address,
			&c.Percent,
		)
		if err != nil {
			return nil, fmt.Errorf("query and add commission failed: %v", err)
		}

		for i := range products {
			if products[i].Product.ProductId == product_id {
				products[i].Product.NovelliaStandardToken.Commission = append(products[i].Product.NovelliaStandardToken.Commission, c)
			}
		}
	}

	return products, nil
}

// queries attribution information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddAttribution(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.pool.Query(ctx, s.queries[queryAttribution], productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a nvla.Attribution
		var product_id string
		err = rows.Scan(
			&product_id,
			&a.AuthorName,
			&a.Url,
			&a.WorkAttributed,
		)
		if err != nil {
			return nil, fmt.Errorf("query and add attribution failed: %v", err)
		}

		for i := range products {
			if products[i].Product.ProductId == product_id {
				products[i].Attribution = append(products[i].Attribution, a)
			}
		}
	}

	return products, nil
}

// queries remote resource information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddRemoteResource(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.pool.Query(ctx, s.queries[queryRemoteResource], productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r nvla.OffChainResource
		var product_id string
		err = rows.Scan(
			&product_id,
			&r.ResourceId,
			&r.Description,
			&r.Priority,
			&r.Multihash,
			&r.HashSourceType,
			&r.Url,
			&r.ContentType,
		)
		if err != nil {
			return nil, fmt.Errorf("query and add remote resource failed: %v", err)
		}

		for i := range products {
			if products[i].Product.ProductId == product_id {
				products[i].Product.NovelliaStandardToken.Resource = append(products[i].Product.NovelliaStandardToken.Resource, r)
			}
		}
	}

	return products, nil
}

func (s *ServiceImpl) Close(ctx context.Context) {
	s.pool.Close()
}
