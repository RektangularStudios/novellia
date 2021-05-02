package novellia_database

import (
	"fmt"
	"context"
	"io/ioutil"
	"path/filepath"
	"github.com/jackc/pgx/v4"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
)

const (
	queryProductFilename = "query_product.sql"
	queryCommissionFilename = "query_commission.sql"
	queryAttributionFilename = "query_attribution.sql"
	queryRemoteResourceFilename = "query_remote_resource.sql"
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

// reads a text file using the queriesPath as the base path
func (s *ServiceImpl) ReadQueryFile(filename string) (string, error) {
	queryPath := filepath.Join(s.queriesPath, filename)

	bytes, err := ioutil.ReadFile(queryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read query file %s: %v", filename, err)
	}

	return string(bytes), nil
}

// queries product information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddProduct(ctx context.Context, products []nvla.Product) ([]nvla.Product, error) {
	productQuery, err := s.ReadQueryFile(queryProductFilename)
	if err != nil {
		return nil, err
	}

	rows, err := s.conn.Query(ctx, productQuery)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	/*
	for _, field := range rows.FieldDescriptions() {
		fmt.Printf("%+v", string(field.Name[:]))
	}
	*/

	for rows.Next() {
		var p nvla.Product
		var t nvla.NovelliaStandardToken
		p.Product.NovelliaStandardToken = &t
		var dateListed pgtype.Date
		var dateAvailable pgtype.Date

		// TODO: handle case where native token does not exist (i.e. NovelliaProduct)
		err = rows.Scan(
			// product
			&p.Product.ProductId, &p.Product.NovelliaStandardToken.Name,
			&p.Pricing.PriceCurrencyId, &p.Pricing.PriceUnitAmount, &p.Pricing.MaxOrderSize,
			&dateListed, &dateAvailable,
			// organization
			&p.Organization.OrganizationId, &p.Organization.Name,
			// market
			&p.Market.MarketId, &p.Market.Name,
			// native token
			&t.NativeToken.PolicyId, &t.NativeToken.AssetId,
			// product detail
			&t.Copyright,
			&t.Publisher,
			&t.Version,
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
		
		&p.Metadata.DateListed, &p.Metadata.DateAvailable,

		fmt.Println(err)
		fmt.Printf("\nrow: %+v\n%+v\n\n", p, t)

		// add product to slice
		products = append(products, p)
	}

	return products, nil
}

/*
// queries commission information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddCommission(ctx context.Context, products []nvla.Product) ([]nvla.Product, error) {

}

// queries attribution information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddAttribution(ctx context.Context, products []nvla.Product) ([]nvla.Product, error) {

}

// queries remote resource information and adds it to the provided products slice
func (s *ServiceImpl) QueryAndAddRemoteResource(ctx context.Context, products []nvla.Product) ([]nvla.Product, error) {

}
*/

func (s *ServiceImpl) Close(ctx context.Context) {
	s.conn.Close(ctx)
}
