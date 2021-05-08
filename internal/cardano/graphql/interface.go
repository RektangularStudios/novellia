package cardano_graphql

import (
	"context"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type Service interface {
	// Status of the remote GraphQL instance, and the sync percentage
	GetStatus(ctx context.Context) (bool, float32, error)
	// Tip of Cardano, returns lastest block and epoch
	GetTip(ctx context.Context) (int32, int32, error)
	// Assets owned at a payment address
	GetAssets(ctx context.Context, paymentAddress string) ([]nvla.Token, error)
}
