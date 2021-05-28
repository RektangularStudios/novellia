package novellia_database

import (
	"context"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type Service interface {
	QueryProductIDs(ctx context.Context, organizationId string, marketId string) ([]nvla.ProductListElement, error)
	QueryAndAddProduct(ctx context.Context, productElements []nvla.ProductListElement) ([]nvla.Product, error)
	QueryAndAddCommission(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddAttribution(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddRemoteResource(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)	
	QueryAndAddProductModified(ctx context.Context, productIDs []string, products []nvla.Product) ([]nvla.Product, error)
	Close(ctx context.Context)
}
