package api

import (
	"fmt"
	"context"
	"errors"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
	cardano_graphql "github.com/RektangularStudios/novellia/internal/cardano/graphql"
)

type ApiService struct{
	cardanoGraphQLService cardano_graphql.Service
}

// NewApiService creates an api service
func NewApiService(cardanoGraphQLService cardano_graphql.Service) nvla.DefaultApiServicer {
	return &ApiService {
		cardanoGraphQLService: cardanoGraphQLService,
	}
}

// Gets an order by id
func (s *MockedApiService) GetOrders(ctx context.Context, productId string) (nvla.ImplResponse, error) {
	order := nvla.Order{
		Products: []nvla.OrderProducts{
			{
				ProductId: "PROD-01D78XYFJ1PRM1WPBAOU8JQMNV",
				Quantity 4,
			},
			{
				ProductId: "PROD-01D78XYFJ1PRM1WPBCBT3VHMNV",
				Quantity 2,
			},
		}
		Customer: nvla.OrderCustomer{
			DeliveryAddress: "addr1q80u75kavwd5sc7j52x0k8nrqd46540vcjgsvl4fhxjqqs60vcjwf9llp7rv006f0dqyffltyyyzpzl9vct4mp7wjdaspwq39a",
		}
		Payment: nvla.OrderPayment{
			PaymentAddress: "addr1q80u75kavwd5sc7j52x0k8nrqd46540vcjgsvl4fhxjqqs60vcjwf9llp7rv006f0dqyffltyyyzpzl9vct4mp7wjdaspwq39a",
			PriceCurrencyId "ada",
			PriceAmount 20,
			Status string `json:"status"`,
		}
		OrderId: "ORDER-01D78XYFJ1PRM1WPBCBT3VHMNV",
	}

	return nvla.Response(200, order), nil
}

// Creates an order and returns the order_id
func (s *MockedApiService) PostOrders(context.Context, Order) (ImplResponse, error) {
	orderCreated := nvla.OrderCreated{
		OrderId: "ORDER-01D78XYFJ1PRM1WPBCBT3VHMNV",
	}

	return nvla.Response(200, orderCreated), nil
}

// Gets listed products
func (s *ApiService) GetProducts(ctx context.Context, marketId string, organizationId string, productId string) (nvla.ImplResponse, error) {
	// TODO: unmock, use query params

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
func (s *ApiService) GetStatus(ctx context.Context) (nvla.ImplResponse, error) {
	initialized, syncPercentage, err := s.cardanoGraphQLService.GetStatus(ctx)
	
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	resp := nvla.Status{
		Cardano: nvla.StatusCardano{
			Initialized: initialized,
			SyncPercentage: syncPercentage,
		},
		// TODO: read this value from somewhere
		Maintenance: false,
	}
	return nvla.Response(200, resp), nil
}

// Cardano chain tip information
func (s *ApiService) GetCardanoTip(ctx context.Context) (nvla.ImplResponse, error) {
	blockNumber, epochNumber, err := s.cardanoGraphQLService.GetTip(ctx)
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	resp := nvla.CardanoTip{
		Block: blockNumber,
		Epoch: epochNumber,
	}
	return nvla.Response(200, resp), nil
}

// Lists assets owned by a wallet
func (s *ApiService) GetWallet(ctx context.Context, walletAddress string) (nvla.ImplResponse, error) {
	tokens, err := s.cardanoGraphQLService.GetAssets(ctx, walletAddress)
	if err != nil {
		return nvla.Response(500, nil), err
	}
	return nvla.Response(200, tokens), nil
}

// GetWorkflowMinterNvla -
func (s *ApiService) GetWorkflowMinterNvla(ctx context.Context) (nvla.ImplResponse, error) {
	// TODO - update GetWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response nvla.Response(200, WorkflowInformation{}) or use other options such as http.Ok ...
	//return nvla.Response(200, WorkflowInformation{}), nil

	return nvla.Response(http.StatusNotImplemented, nil), errors.New("GetWorkflowMinterNvla method not implemented")
}

// PostCardanoTransaction -
func (s *ApiService) PostCardanoTransaction(ctx context.Context, cardanoTransaction nvla.CardanoTransaction) (nvla.ImplResponse, error) {
	// TODO - update PostCardanoTransaction with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response nvla.Response(200, {}) or use other options such as http.Ok ...
	//return nvla.Response(200, nil),nil

	//TODO: Uncomment the next line to return response nvla.Response(400, {}) or use other options such as http.Ok ...
	//return nvla.Response(400, nil),nil

	return nvla.Response(http.StatusNotImplemented, nil), errors.New("PostCardanoTransaction method not implemented")
}

// PostWorkflowMinterNvla -
func (s *ApiService) PostWorkflowMinterNvla(ctx context.Context, minterInfo nvla.MinterInfo) (nvla.ImplResponse, error) {
	// TODO - update PostWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response nvla.Response(200, CardanoTransaction{}) or use other options such as http.Ok ...
	//return nvla.Response(200, CardanoTransaction{}), nil

	//TODO: Uncomment the next line to return response nvla.Response(400, {}) or use other options such as http.Ok ...
	//return nvla.Response(400, nil),nil

	return nvla.Response(http.StatusNotImplemented, nil), errors.New("PostWorkflowMinterNvla method not implemented")
}
