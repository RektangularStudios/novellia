package api

import (
	"fmt"
	"context"
	"errors"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	cardano_graphql "github.com/RektangularStudios/novellia/internal/cardano/graphql"
	"github.com/RektangularStudios/novellia/internal/novellia_database"
)

type ApiService struct{
	cardanoGraphQLService cardano_graphql.Service
	novelliaDatabaseService novellia_database.Service
}

// NewApiService creates an api service
func NewApiService(
	cardanoGraphQLService cardano_graphql.Service,
	novelliaDatabaseService novellia_database.Service,
) nvla.DefaultApiServicer {
	return &ApiService {
		cardanoGraphQLService: cardanoGraphQLService,
		novelliaDatabaseService: novelliaDatabaseService,
	}
}

// Gets an order by id
func (s *ApiService) GetOrders(ctx context.Context, productId string) (nvla.ImplResponse, error) {
	return nvla.Response(http.StatusNotImplemented, nil), errors.New("GetOrders method not implemented")
}

// Creates an order and returns the order_id
func (s *ApiService) PostOrders(context.Context, nvla.Order) (nvla.ImplResponse, error) {
	return nvla.Response(http.StatusNotImplemented, nil), errors.New("PostOrders method not implemented")
}

// Gets list of products
func (s *ApiService) GetProducts(ctx context.Context, marketId string, organizationId string) (nvla.ImplResponse, error) {
	productIDs, err := s.novelliaDatabaseService.QueryProductIDs(ctx, organizationId, marketId)
	if err != nil {
		err = fmt.Errorf("get products failed at product ID query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	
	productsList := nvla.ProductsList{
		ProductId: productIDs,
	}
	return nvla.Response(200, productsList), nil
}


// Post for list of products details
func (s *ApiService) PostProducts(ctx context.Context, productsList nvla.ProductsList) (nvla.ImplResponse, error) {
	products, err := s.novelliaDatabaseService.QueryAndAddProduct(ctx, productsList.ProductId)
	if err != nil {
		err = fmt.Errorf("post products failed at product query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddCommission(ctx, productsList.ProductId, products)
	if err != nil {
		err = fmt.Errorf("post products failed at commission query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddAttribution(ctx, productsList.ProductId, products)
	if err != nil {
		err = fmt.Errorf("post products failed at attribution query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddRemoteResource(ctx, productsList.ProductId, products)
	if err != nil {
		err = fmt.Errorf("post products failed at remote resource query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	return nvla.Response(200, products), nil
}

// Availability information about service availability
func (s *ApiService) GetStatus(ctx context.Context) (nvla.ImplResponse, error) {
	initialized, syncPercentage, err := s.cardanoGraphQLService.GetStatus(ctx)
	
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	status := "UP"
	if !initialized || syncPercentage != 1 {
		status = "Cardano service not ready"
	}

	resp := nvla.Status{
		Cardano: nvla.StatusCardano{
			Initialized: initialized,
			SyncPercentage: syncPercentage,
		},
		// TODO: read this value from somewhere
		Maintenance: false,
		Status: status,
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
