package api

import (
	"context"
	"errors"
	"net/http"
	"time"
	"fmt"

	"github.com/RektangularStudios/novellia/internal/constants"
	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type MockedApiService struct{}

// MockedNewApiService creates an api service
func NewMockedApiService() nvla.DefaultApiServicer {
	return &MockedApiService {}
}

// Gets list of products
func (s *MockedApiService) GetProducts(ctx context.Context, marketId string, organizationId string) (nvla.ImplResponse, error) {
	tokenProduct := s.GetMockNovelliaStandardTokenProduct()
	product := s.GetMockNovelliaProduct()

	productsList := nvla.ProductsList{
		Products: []nvla.ProductListElement{
			nvla.ProductListElement{
				ProductId: tokenProduct.Product.ProductId,
				NativeTokenId: fmt.Sprintf("%s.%s", tokenProduct.Product.NovelliaStandardToken.NativeToken.PolicyId, tokenProduct.Product.NovelliaStandardToken.NativeToken.AssetId),
				Modified: time.Now().UTC().Format(constants.ISO8601DateFormat),
			},
			nvla.ProductListElement{
				ProductId: product.Product.ProductId,
				NativeTokenId: "",
				Modified: time.Now().UTC().Format(constants.ISO8601DateFormat),
			},
		},
	}

	return nvla.Response(200, productsList), nil
}

// Post for list of products details
func (s *MockedApiService) PostProducts(ctx context.Context, productsList nvla.ProductsList) (nvla.ImplResponse, error) {
	novelliaStandardTokenProduct := s.GetMockNovelliaStandardTokenProduct()
	novelliaProduct := s.GetMockNovelliaProduct()

	return nvla.Response(200, []nvla.Product{novelliaStandardTokenProduct, novelliaProduct}), nil
}

// Availability information about service availability
func (s *MockedApiService) GetStatus(ctx context.Context) (nvla.ImplResponse, error) {
	resp := nvla.Status{
		Cardano: nvla.StatusCardano{
			Initialized: true,
			SyncPercentage: 100,
		},
		Maintenance: false,
		Status: "UP",
	}

	return nvla.Response(200, resp), nil
}

// Cardano chain tip information
func (s *MockedApiService) GetCardanoTip(ctx context.Context) (nvla.ImplResponse, error) {
	resp := nvla.CardanoTip{
		Block: 5622050,
		Epoch: 261,
	}

	return nvla.Response(200, resp), nil
}

// Lists assets owned by a wallet
func (s *MockedApiService) PostWallet(ctx context.Context, wallet nvla.Wallet) (nvla.ImplResponse, error) {
	tokens := []nvla.Token{
		nvla.Token{
			NativeTokenId: "0xOccultaNovellia.IscaraTheTenThousandGuns",
			Amount: 2400,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
		},
		nvla.Token{
			NativeTokenId: "0xOccultaNovellia.Draculi",
			Amount: 500,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
		},
		nvla.Token{
			NativeTokenId: "0xOccultaNovellia.Voyin",
			Amount: 0,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
		},
	}

	return nvla.Response(200, tokens), nil
}

// PostTokens
func (s *MockedApiService) PostTokens(ctx context.Context, tokenSearch nvla.TokenSearch) (nvla.ImplResponse, error) {
	tokens := []nvla.Token{
		nvla.Token{
			NativeTokenId: "0xOccultaNovellia.AlphaDraculi",
			InitialMintTxMetadata: "{\"version\": 1, \"copyright\": \"Copyright Rektangular Studios Inc.; all rights reserved\", \"extension\": [\"novellia_1\"], \"publisher\": [\"https://rektangularstudios.com\"], \"d27dadf7c5f24bfe9e377927c2d811d63d19222e1a53bb50cbb51772\": {\"AlphaDraculi\": {\"id\": 3, \"name\": \"Alpha Draculi\", \"tags\": [\"Collectible Character\"], \"image\": \"ipfs://QmTUkcUb2EHSw7Ptd52zeuKwgf8G4iVMczXJdMz8XN3dzd\", \"resource\": [{\"url\": [\"https://api.rektangularstudios.com/static/r0tngydjh7x0/nvla.json\", \"ipfs://QmdEN1b9xi5R44tssj5CojMBCwXYTh5P6pHKV6YcPogCXZ\", \"sia://AAAG1QULaSR9sguw63RnlyGGDYsygLW8rbHlG-_kcf9IYQ\"], \"priority\": 0, \"multihash\": \"QmdEN1b9xi5R44tssj5CojMBCwXYTh5P6pHKV6YcPogCXZ\", \"description\": \"Off-chain Novellia extended metadata\", \"resource_id\": \"Novellia\", \"content_type\": \"application/json\", \"hash_source_type\": \"ipfs\"}], \"description\": {\"long\": \"A character for the surreal horror multiverse Occulta Novellia\", \"short\": \"Occulta Novellia Character\"}}}}",
		},
	}

	return nvla.Response(200, tokens), nil
}

// GetWorkflowMinterNvla -
func (s *MockedApiService) GetWorkflowMinterNvla(ctx context.Context) (nvla.ImplResponse, error) {
	return nvla.Response(http.StatusNotImplemented, nil), errors.New("GetWorkflowMinterNvla method not implemented")
}

// PostCardanoTransaction -
func (s *MockedApiService) PostCardanoTransaction(ctx context.Context, cardanoTransaction nvla.CardanoTransaction) (nvla.ImplResponse, error) {
	return nvla.Response(http.StatusNotImplemented, nil), errors.New("PostCardanoTransaction method not implemented")
}

// PostWorkflowMinterNvla -
func (s *MockedApiService) PostWorkflowMinterNvla(ctx context.Context, minterInfo nvla.MinterInfo) (nvla.ImplResponse, error) {
	return nvla.Response(http.StatusNotImplemented, nil), errors.New("PostWorkflowMinterNvla method not implemented")
}
