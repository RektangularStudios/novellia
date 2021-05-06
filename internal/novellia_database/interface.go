package novellia_database

import (
	"context"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
)

type Service interface {
	QueryAndAddProduct(ctx context.Context, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddCommission(ctx context.Context, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddAttribution(ctx context.Context, products []nvla.Product) ([]nvla.Product, error)
	QueryAndAddRemoteResource(ctx context.Context, products []nvla.Product) ([]nvla.Product, error)	
	Close(ctx context.Context)
}
