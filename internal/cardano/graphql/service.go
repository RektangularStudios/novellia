package cardano_graphql

import (
	"fmt"
	"context"
	"strconv"
	"github.com/shurcooL/graphql"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type AddressInfo struct {
	Address string `json:"address"`
	Base16 string `json:"base16"`
	Type string `json:"type"`
	Encoding string `json:"encoding"`
	Era string `json:"era"`
}

type ServiceImpl struct {
	graphQLClient *graphql.Client
}

func New(graphQLClient *graphql.Client) (*ServiceImpl) {
	return &ServiceImpl {
		graphQLClient: graphQLClient,
	}
}

// Returns the initialization status of the remote GraphQL instance, and the sync percentage
func (s *ServiceImpl) GetStatus(ctx context.Context) (bool, float32, error) {
	var query struct {
		CardanoDbMeta struct {
			Initialized graphql.Boolean
			SyncPercentage graphql.Float
		}
	}

	err := s.graphQLClient.Query(ctx, &query, nil)
	if err != nil {
		return false, 0, err
	}

	return bool(query.CardanoDbMeta.Initialized), float32(query.CardanoDbMeta.SyncPercentage), nil
}

// Tip of Cardano, returns lastest block and epoch
func (s *ServiceImpl) GetTip(ctx context.Context) (int32, int32, error) {
	var query struct {
		Cardano struct {
			Tip struct {
				Number graphql.Int
				EpochNo graphql.Int
			}
		}
	}

	err := s.graphQLClient.Query(ctx, &query, nil)
	if err != nil {
		return 0, 0, err
	}

	return int32(query.Cardano.Tip.Number), int32(query.Cardano.Tip.EpochNo), nil
}

func (s *ServiceImpl) GetAssets(ctx context.Context, paymentAddress string) ([]nvla.Token, error) {
	// query assets at latest block
	blockNumber, _, err := s.GetTip(ctx)
	if err != nil {
		return nil, err
	}

	queryParams := map[string]interface{}{
		"addresses":  []graphql.String{graphql.String(paymentAddress)},
		"atBlock": 		graphql.Int(blockNumber),
	}
	
	var query struct {
		PaymentAddresses []struct {
			Summary struct {
				AssetBalances []struct {
					Asset struct{
						AssetID graphql.String
						Description graphql.String
						Name graphql.String
					}
					Quantity graphql.String
				}
				UtxosCount graphql.Int
			} `graphql:"summary(atBlock: $atBlock)"`
		} `graphql:"paymentAddresses(addresses: $addresses)"`
	}

	err = s.graphQLClient.Query(ctx, &query, queryParams)
	if err != nil {
		return nil, err
	}

	if len(query.PaymentAddresses) != 1 {
		return nil, fmt.Errorf("expected a single payment address with assets: %v", err)
	}
	tokens := []nvla.Token{}
	for _, assetBalance := range query.PaymentAddresses[0].Summary.AssetBalances {
		amount, err := strconv.ParseInt(string(assetBalance.Quantity), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("could not convert balance to integer: %v", err)
		}

		t := nvla.Token{
			Amount: int32(amount),
			Name: string(assetBalance.Asset.Name),
			AssetId: string(assetBalance.Asset.AssetID),
			Description: string(assetBalance.Asset.Description),
		}

		tokens = append(tokens, t)
	}

	return tokens, nil
}

func (s *ServiceImpl) GetAddressType(address string) (string, error) {
	out, err := exec.Command("cardano-cli", "address", "info",
		"--address", address,
	).Output()
	if err != nil {
		return fmt.Errorf("failed to validate address (cmd): %v", err)
	}
	if strings.Contains(string(out), "Invalid") {
		return fmt.Errorf("address is invalid: %s, output: %s", address, string(out))
	}

	var info AddressInfo
	err = json.Unmarshal(out, &info)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal address info JSON: %v", err)
	}

	return info.Type, nil
}
