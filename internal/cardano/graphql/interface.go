package cardano_graphql

import (
	"context"
)

type Service interface {
	// Returns the initialization status of the remote GraphQL instance, and the sync percentage
	Initialized(ctx context.Context) (bool, float64, error)
}
