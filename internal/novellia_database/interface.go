package novellia_database

import (
	"context"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type Service interface {
	PrepareQueries(ctx context.Context) error
	QueryProductIDs(ctx context.Context, organizationId string, marketId string) ([]string, error)
	QueryAndAddProduct(ctx context.Context, productIDs []string) ([]nvla.Product, error)
	QueryAndAddCommission(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddAttribution(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddRemoteResource(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)	
	Close(ctx context.Context)
}
