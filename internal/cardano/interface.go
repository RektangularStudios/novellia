package cardano

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
	GetAssets(ctx context.Context, wallet nvla.Wallet) ([]nvla.Token, error)
	// Gets an address' type
	GetAddressType(address string) (string, error)
	// Add 721 onchain metadata to list of tokens
	Add721Metadata(ctx context.Context, tokens []nvla.Token) ([]nvla.Token, error)
	// Queries the stake key associated with a payment address (if it exists)
	QueryStakeKeyFromPaymentAddress(ctx context.Context, paymentAddress string) (string, error)
	// Queries payment addresses from stake key
	QueryPaymentAddressesesFromStakeKey(ctx context.Context, stakeKey string) ([]string, error)
	// Queries ADA held in a list of payment addresses
	QueryADABalance(ctx context.Context, paymentAddresses []string, tokens []nvla.Token) ([]nvla.Token, error)
	// Queries tokens held in a list of payment addresses
	QueryTokenBalance(ctx context.Context, paymentAddresses []string, tokens []nvla.Token) ([]nvla.Token, error)
	// Safely closes the DB connection
	Close(ctx context.Context)
}
