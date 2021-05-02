package novellia_database

import (
	"context"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
)

type Service interface {
	ReadQueryFile(path string) (string, error)
	QueryAndAddProduct(ctx context.Context, products []nvla.Product) ([]nvla.Product, error)
	Close(ctx context.Context)
}
