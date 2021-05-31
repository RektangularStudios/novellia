package cardano

import (
	"fmt"
	"context"
	"strconv"
	"os/exec"
	"encoding/json"
	"strings"
	"io/ioutil"
	"path/filepath"
	"encoding/hex"

	"github.com/shurcooL/graphql"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/RektangularStudios/novellia/internal/config"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

const (
	graphQLRetries = 10
)

const (
	queryMetadata = "queryMetadata"
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
	queriesPath string
	pool *pgxpool.Pool
	queries map[string]string
}

func New(ctx context.Context) (*ServiceImpl, error) {
	cfg, err := config.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get config")
	}

	graphQLHostString := fmt.Sprintf("%s:%s", cfg.CardanoGraphQL.Host, cfg.CardanoGraphQL.Port)
	graphQLClient := graphql.NewClient(graphQLHostString, nil)

	// url like "postgresql://username:password@localhost:5432/database_name"
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		cfg.Postgres.Username,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.CardanoDatabase,
	)
	pool, err := pgxpool.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Postgres: %v", err)
	}
	
	service := ServiceImpl {
		graphQLClient: graphQLClient,
		pool: pool,
		queriesPath: cfg.Postgres.QueriesPath,
	}
	err = service.loadQueries(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load queries")
	}

	return &service, nil
}

func (s *ServiceImpl) loadQueries(ctx context.Context) error {
	queryFiles := map[string]string {
		queryMetadata: "query_metadata.sql",
	}

	queries := make(map[string]string)
	for name, filename := range queryFiles {
		fmt.Printf("Loading SQL %s\n", filename)

		query, err := s.readQueryFile(filename)
		if err != nil {
			return err
		}

		queries[name] = query
	}
	s.queries = queries

	fmt.Printf("SQL has been loaded\n")
	return nil
}

// reads a text file using the queriesPath as the base path
func (s *ServiceImpl) readQueryFile(filename string) (string, error) {
	queryPath := filepath.Join(s.queriesPath, filename)

	bytes, err := ioutil.ReadFile(queryPath)
	if err != nil {
		return "", fmt.Errorf("failed to read query file %s: %v", filename, err)
	}

	return string(bytes), nil
}

// Returns the initialization status of the remote GraphQL instance, and the sync percentage
func (s *ServiceImpl) GetStatus(ctx context.Context) (bool, float32, error) {
	var query struct {
		CardanoDbMeta struct {
			Initialized graphql.Boolean
			SyncPercentage graphql.Float
		}
	}

	// TODO: remove this loop once cardano-graphql is less buggy
	for i := 0; i < graphQLRetries; i++ {
		err := s.graphQLClient.Query(ctx, &query, nil)
		if err != nil && i == graphQLRetries - 1 {
			return false, 0, err
		}
		if err != nil {
			continue
		}
		break
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
		break
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
					Asset struct {
						PolicyID graphql.String
						Description graphql.String
						Name graphql.String
						AssetName graphql.String
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
		break
	}

	for _, paymentAddress := range query.PaymentAddresses {
		for _, assetBalance := range paymentAddress.Summary.AssetBalances {
			amount, err := strconv.ParseUint(string(assetBalance.Quantity), 10, 64)
			if err != nil {
				return nil, fmt.Errorf("could not convert balance to integer: %v", err)
			}

			nativeTokenID := "ada"
			if string(assetBalance.Asset.AssetName) != "ada" {
				decodedAssetID, err := hex.DecodeString(string(assetBalance.Asset.AssetName))
				if err != nil {
					return nil, err
				}
				nativeTokenID = fmt.Sprintf("%s.%s", string(assetBalance.Asset.PolicyID), string(decodedAssetID))
			}
			t := nvla.Token{
				Amount: uint64(amount),
				Name: string(assetBalance.Asset.Name),
				NativeTokenId: nativeTokenID,
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

func (s *ServiceImpl) query721Metadata(ctx context.Context, nativeTokens []nvla.NativeToken) (map[string]string, error) {
	metadataJSON := map[string]string{}

	policyIDs := [][]byte{}
	assetIDs := [][]byte{}
	for _, n := range nativeTokens {
		decodedPolicyID, err := hex.DecodeString(n.PolicyId)
		if err != nil {
			return nil, err
		}
		policyIDs = append(policyIDs, decodedPolicyID)

		assetIDs = append(assetIDs, []byte(n.AssetId))
	}

	rows, err := s.pool.Query(ctx, s.queries[queryMetadata], policyIDs, assetIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var policyID, assetID []byte
		var j string
		err = rows.Scan(
			&policyID,
			&assetID,
			&j,
		)
		if err != nil {
			return nil, fmt.Errorf("query 721 metadata failed: %v", err)
		}

		nativeTokenID := fmt.Sprintf("%x.%s", policyID, string(assetID))
		metadataJSON[nativeTokenID] = j
	}

	return metadataJSON, nil
}

func (s *ServiceImpl) Add721Metadata(ctx context.Context, tokens []nvla.Token) ([]nvla.Token, error) {
	nativeTokens := []nvla.NativeToken{}
	for _, t := range tokens {
		if t.NativeTokenId == "ada" {
			continue
		}
		f := strings.Split(t.NativeTokenId, ".")
		if len(f) != 2 {
			return nil, fmt.Errorf("failed to split native token ID into policy and asset IDs")
		}

		policyID := f[0]
		assetID := f[1]
		nativeTokens = append(nativeTokens, nvla.NativeToken{
			PolicyId: policyID,
			AssetId: assetID,
		})
	}

	metadataJSON, err := s.query721Metadata(ctx, nativeTokens)
	if err != nil {
		return nil, err
	}

	for i := range tokens {
		if j, ok := metadataJSON[tokens[i].NativeTokenId]; ok {
			tokens[i].InitialMintTxMetadata = j
		}
	}

	return tokens, nil
}

func (s *ServiceImpl) Close(ctx context.Context) {
	s.pool.Close()
}
