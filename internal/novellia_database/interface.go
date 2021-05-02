package novellia_database

import (
	"context"
)

type Service interface {
	ListProductAttribution() error
	Close(ctx context.Context)
}
