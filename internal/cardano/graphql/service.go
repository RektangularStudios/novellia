package cardano_graphql

import (
	"fmt"
	"context"
	"strconv"
	"os/exec"
	"encoding/json"
	"strings"
	"github.com/shurcooL/graphql"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

const (
	graphQLRetries = 10
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

	// TODO: remove this loop once cardano-graphql is less buggy
	for i := 0; i < graphQLRetries; i++ {
		err := s.graphQLClient.Query(ctx, &query, nil)
		if err != nil && i == graphQLRetries - 1 {
			return 0, 0, err
		}
		if err != nil {
			continue
		}
	}


	return int32(query.Cardano.Tip.Number), int32(query.Cardano.Tip.EpochNo), nil
}

func (s *ServiceImpl) categorizeWalletIdentifiers(wallet nvla.Wallet) ([]string, []string, error) {
	paymentAddresses := []string{}
	stakeAddresses := []string{}

	for _, addr := range wallet.CardanoIdentifiers {
		t, err := s.GetAddressType(addr)
		if err != nil {
			return nil, nil, err
		}

		switch t {
		case "payment":
			paymentAddresses = append(paymentAddresses, addr)
		case "stake":
			stakeAddresses = append(stakeAddresses, addr)
		}
	}

	return paymentAddresses, stakeAddresses, nil
}

func (s *ServiceImpl) getPaymentAddressesFromStakeAddress(ctx context.Context, stakeAddress string) ([]string, error) {
	paymentAddresses := []string{}
	
	return paymentAddresses, nil
}

func (s *ServiceImpl) getAssetsFromPaymentAddresses(ctx context.Context, paymentAddresses []string) ([]nvla.Token, error) {
	// query assets at latest block
	blockNumber, _, err := s.GetTip(ctx)
	if err != nil {
		return nil, err
	}

	tokens := []nvla.Token{}
	if len(paymentAddresses) == 0 {
		return tokens, nil
	}

	graphQLPaymentAddresses := []graphql.String{}
	for _, p := range paymentAddresses {
		graphQLPaymentAddresses = append(graphQLPaymentAddresses, graphql.String(p))
	}

	queryParams := map[string]interface{}{
		"addresses":  graphQLPaymentAddresses,
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

	// TODO: remove this loop once cardano-graphql is less buggy
	for i := 0; i < graphQLRetries; i++ {
		err = s.graphQLClient.Query(ctx, &query, queryParams)
		if err != nil && i == graphQLRetries - 1 {
			return nil, err
		}
		if err != nil {
			continue
		}
	}

	for _, paymentAddress := range query.PaymentAddresses {
		for _, assetBalance := range paymentAddress.Summary.AssetBalances {
			amount, err := strconv.ParseInt(string(assetBalance.Quantity), 10, 32)
			if err != nil {
				return nil, fmt.Errorf("could not convert balance to integer: %v", err)
			}

			t := nvla.Token{
				Amount: uint64(amount),
				Name: string(assetBalance.Asset.Name),
				NativeTokenId: string(assetBalance.Asset.AssetID),
				Description: string(assetBalance.Asset.Description),
			}

			tokens = append(tokens, t)
		}
	}

	return tokens, nil
}

func (s *ServiceImpl) GetAssets(ctx context.Context, wallet nvla.Wallet) ([]nvla.Token, error) {
	paymentAddresses, stakeAddresses, err := s.categorizeWalletIdentifiers(wallet)
	if err != nil {
		return nil, err
	}	

	for _, stakeAddr := range(stakeAddresses) {
		p, err := s.getPaymentAddressesFromStakeAddress(ctx, stakeAddr)
		if err != nil {
			return nil, err
		}
		paymentAddresses = append(paymentAddresses, p...)
	}

	tokens, err := s.getAssetsFromPaymentAddresses(ctx, paymentAddresses)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (s *ServiceImpl) GetAddressType(address string) (string, error) {
	out, err := exec.Command("cardano-cli", "address", "info",
		"--address", address,
	).Output()
	if err != nil {
		return "", fmt.Errorf("failed to validate address (cmd): %v", err)
	}
	if strings.Contains(string(out), "Invalid") {
		return "", fmt.Errorf("address is invalid: %s, output: %s", address, string(out))
	}

	var info AddressInfo
	err = json.Unmarshal(out, &info)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal address info JSON: %v", err)
	}

	if info.Era != "shelley" {
		return "", fmt.Errorf("not a shelley address: %s", address)
	}

	return info.Type, nil
}
