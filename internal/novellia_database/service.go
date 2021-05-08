package novellia_database

import (
	"fmt"
	"context"
	"io/ioutil"
	"path/filepath"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgtype"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
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
	conn *pgx.Conn
}

// creates a new ServiceImpl, connecting to Postgres
func New(ctx context.Context, username, password, host, database_name string, queriesPath string) (*ServiceImpl, error) {
	// url like "postgresql://username:password@localhost:5432/database_name"
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s", username, password, host, database_name)
	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Postgres: %v", err)
	}

	return &ServiceImpl {
		conn: conn,
		queriesPath: queriesPath,
	}, nil
}

func (s *ServiceImpl) PrepareQueries(ctx context.Context) error {
	queries := map[string]string {
		queryProductID: "query_product_id.sql",
		queryProduct: "query_product.sql",
		queryCommission: "query_commission.sql",
		queryAttribution: "query_attribution.sql",
		queryRemoteResource: "query_remote_resource.sql",
	}
	
	for name, filename := range queries {
		fmt.Printf("Preparing SQL %s\n", filename)

		query, err := s.readQueryFile(filename)
		if err != nil {
			return err
		}
		_, err = s.conn.Prepare(ctx, name, query)
		if err != nil {
			return err
		}
	}
	fmt.Printf("SQL preparations done\n")
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
func (s *ServiceImpl) QueryProductIDs(ctx context.Context, organizationId string, marketId string) ([]string, error) {
	rows, err := s.conn.Query(ctx, queryProductID, organizationId, marketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productIDs := []string{}
	for rows.Next() {
		var productID string;

		err = rows.Scan(
			&productID,
		)
		if err != nil {
			return nil, fmt.Errorf("query product IDs failed: %v", err)
		}

		// add product ID to slice
		productIDs = append(productIDs, productID)
	}

	return productIDs, nil
}

// queries product information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddProduct(ctx context.Context, productIDs []string) ([]nvla.Product, error) {
	rows, err := s.conn.Query(ctx, queryProduct, productIDs)
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

		// TODO: handle case where native token does not exist (i.e. NovelliaProduct)
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
		p.Metadata.DateListed = dateListed.Time.String()
		p.Metadata.DateAvailable = dateAvailable.Time.String()

		// add product to slice
		products = append(products, p)
	}

	return products, nil
}

// queries commission information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddCommission(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.conn.Query(ctx, queryCommission, productIDs)
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

		for _, p := range products {
			if p.Product.ProductId == product_id {
				p.Product.NovelliaStandardToken.Commission = append(p.Product.NovelliaStandardToken.Commission, c)
			}
		}
	}

	return products, nil
}

// queries attribution information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddAttribution(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.conn.Query(ctx, queryAttribution, productIDs)
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

		for _, p := range products {
			if p.Product.ProductId == product_id {
				p.Attribution = append(p.Attribution, a)
			}
		}
	}

	return products, nil
}

// queries remote resource information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddRemoteResource(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error) {
	rows, err := s.conn.Query(ctx, queryRemoteResource, productIDs)
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

		for _, p := range products {
			if p.Product.ProductId == product_id {
				p.Product.NovelliaStandardToken.Resource = append(p.Product.NovelliaStandardToken.Resource, r)
			}
		}
	}

	return products, nil
}

func (s *ServiceImpl) Close(ctx context.Context) {
	s.conn.Close(ctx)
}
