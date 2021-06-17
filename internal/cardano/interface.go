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
	// Gets an address' type and returns the base16 encoding
	GetAddressType(address string) (string, string, error)
	// Gets a stake key embedded in a payment address
	DecodeStakeAddressFromBase16(paymentAddressBase16 string) (string, error)
	// Add 721 onchain metadata to list of tokens
	Add721Metadata(ctx context.Context, tokens []nvla.Token) ([]nvla.Token, error)
	// Queries payment addresses from stake key
	QueryPaymentAddressesesFromStakeKey(ctx context.Context, stakeKey string) ([]string, error)
	// Queries ADA held in a list of payment addresses
	QueryADABalance(ctx context.Context, paymentAddresses []string, tokens []nvla.Token) ([]nvla.Token, error)
	// Queries tokens held in a list of payment addresses
	QueryTokenBalance(ctx context.Context, paymentAddresses []string, tokens []nvla.Token) ([]nvla.Token, error)
	// Query tokens on Cardano from search identifiers
	QueryTokens(ctx context.Context, search nvla.TokenSearch) ([]nvla.Token, error)
	// Safely closes the DB connection
	Close(ctx context.Context)
}
