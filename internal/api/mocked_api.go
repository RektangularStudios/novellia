package api

import (
	"context"
	"errors"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
)

type MockedApiService struct{}

// MockedNewApiService creates an api service
func NewMockedApiService() nvla.DefaultApiServicer {
	return &MockedApiService {}
}

// Gets recent orders
func (s *MockedApiService) GetOrders(ctx context.Context, productId string, marketId string, organizationId string, count string) (nvla.ImplResponse, error) {
	orders := []nvla.Order{
		nvla.Order{
			ProductId: 2,
			OrderSize: 5,
			CurrencyPolicyId: "ada",
			UnitPrice: 10,
		},
		nvla.Order{
			ProductId: 2,
			OrderSize: 2,
			CurrencyPolicyId: "ada",
			UnitPrice: 10,
		},
		nvla.Order{
			ProductId: 2,
			OrderSize: 3,
			CurrencyPolicyId: "ada",
			UnitPrice: 10,
		},
	}

	return nvla.Response(200, orders), nil
}

// Gets listed products
func (s *MockedApiService) GetProducts(ctx context.Context, marketId string, organizationId string, productId string) (nvla.ImplResponse, error) {
	products := []nvla.Product{
		{
			Asset: nvla.ProductAsset{
				Urls: []string{
					"https://siasky.net/AACQHdh6YsQfFshLfHjhjsQNCwKHbnaJ2CryhZRAs4HqNQ",
					"https://api.rektangularstudios.com/ipfs/QmXR2TGqndCHmu4utjpzrSwaGtUVMvrjBQ8eKwpRAnPnTh",
					"https://api.rektangularstudios.com/static/cards/iscara_the_ten_thousand_guns/iscara_the_ten_thousand_guns.zip",
				},
				Loader: "occulta_novellia_character",
			},
			Pricing: nvla.ProductPricing{
				CurrencyPolicyId: "ada",
				UnitPrice: 20,
				MaxOrderSize: 5,
				PurchaseAddress: "0xDraculiDepositAddress",
				DateAvailable: "2021-05-15",
			},
			Organization: nvla.ProductOrganization{
				Name: "Rektangular Studios",
				OrganizationId: 1,
			},
			Market: nvla.ProductMarket{
				Name: "Occulta Novellia",
				MarketId: 1,
			},
			Stock: nvla.ProductStock{
				Available: 2400,
				TotalSupply: 2500,
			},
			Metadata: nvla.ProductMetadata{
				Tags: []string{"Game Character"},
				DateListed: "2021-05-01",
			},
			Artist: nvla.ProductArtist{
				Name: "ArtistName",
				Urls: []string{
					"https://www.artstation.com/",
					"https://www.deviantart.com/",
				},
			},
			Product: nvla.ProductProduct{
				OverviewImageUrls: []string{
					"https://siasky.net/_Aqfe5wxnzn54sPn2cBPNnwmjQ2te4rXZfVNvUb79QhBWw",
					"https://api.rektangularstudios.com/ipfs/QmQbju6V8vpjKS1AG9vw4mQucYoeiCY2dNSnrsWzfJAZue",
					"https://api.rektangularstudios.com/static/cards/iscara_the_ten_thousand_guns/iscara_the_ten_thousand_guns_card.jpg",
				},
				TokenPolicyId: "0xrandomNumbers.IscaraTheTenThousandGuns",
				DescriptionShort: "Occulta Novellia Character",
				DescriptionLong: "A character token for the upcoming surreal horror game Occulta Novellia. This token will grant access to a playable character.",
				Name: "Iscara the Ten Thousand Guns",
				ProductId: 2,
			},
		},
	}

	return nvla.Response(200, products), nil
}

// Availability information about service availability
func (s *MockedApiService) GetStatus(ctx context.Context) (nvla.ImplResponse, error) {
	resp := nvla.Status{
		Cardano: nvla.StatusCardano{
			Initialized: true,
			SyncPercentage: 100,
		},
		Maintenance: false,
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
func (s *MockedApiService) GetWallet(ctx context.Context, walletAddress string) (nvla.ImplResponse, error) {
	tokens := []nvla.Token{
		nvla.Token{
			AssetId: "0xOccultaNovellia.IscaraTheTenThousandGuns",
			Amount: 2400,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
		},
		nvla.Token{
			AssetId: "0xOccultaNovellia.Draculi",
			Amount: 500,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
		},
		nvla.Token{
			AssetId: "0xOccultaNovellia.Voyin",
			Amount: 0,
			Name: "OCCLT",
			Description: "Occulta Novellia Character",
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
