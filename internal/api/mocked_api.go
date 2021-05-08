package api

import (
	"context"
	"errors"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
)

type MockedApiService struct{}

// MockedNewApiService creates an api service
func NewMockedApiService() nvla.DefaultApiServicer {
	return &MockedApiService {}
}

// Gets an order by id
func (s *MockedApiService) GetOrders(ctx context.Context, productId string) (nvla.ImplResponse, error) {
	order := nvla.Order{
		Items: []nvla.OrderItems{
			nvla.OrderItems{
				ProductId: "PROD-01D78XYFJ1PRM1WPBAOU8JQMNV",
				Quantity: 4,
			},
			nvla.OrderItems{
				ProductId: "PROD-01D78XYFJ1PRM1WPBCBT3VHMNV",
				Quantity: 2,
			},
		},
		Customer: nvla.OrderCustomer{
			DeliveryAddress: "addr1q80u75kavwd5sc7j52x0k8nrqd46540vcjgsvl4fhxjqqs60vcjwf9llp7rv006f0dqyffltyyyzpzl9vct4mp7wjdaspwq39a",
		},
		Payment: nvla.OrderPayment{
			PaymentAddress: "addr1q80u75kavwd5sc7j52x0k8nrqd46540vcjgsvl4fhxjqqs60vcjwf9llp7rv006f0dqyffltyyyzpzl9vct4mp7wjdaspwq39a",
			PriceCurrencyId: "ada",
			PriceAmount: 20,
			PaymentStatus: "AWAITING_PAYMENT",
		},
		OrderStatus: "AWAITING_PAYMENT",
		Description: "Occulta Novellia Presale Order",
		OrderId: "ORDER-01D78XYFJ1PRM1WPBCBT3VHMNV",
	}

	return nvla.Response(200, order), nil
}

// Creates an order and returns the order_id
func (s *MockedApiService) PostOrders(context.Context, nvla.Order) (nvla.ImplResponse, error) {
	orderCreated := nvla.OrderCreated{
		OrderId: "ORDER-01D78XYFJ1PRM1WPBCBT3VHMNV",
	}

	return nvla.Response(200, orderCreated), nil
}

// Gets list of products
func (s *MockedApiService) GetProducts(ctx context.Context, marketId string, organizationId string) (nvla.ImplResponse, error) {
	productsList := nvla.ProductsList{
		ProductId: []string{
			MockedNovelliaStandardTokenProductId,
			MockedNovelliaProductProductId,
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
