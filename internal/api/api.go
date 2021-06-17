package api

import (
	"fmt"
	"context"
	"errors"
	"net/http"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/novellia/v0"
	"github.com/RektangularStudios/novellia/internal/cardano"
	"github.com/RektangularStudios/novellia/internal/novellia_database"
)

type ApiService struct{
	cardanoService cardano.Service
	novelliaDatabaseService novellia_database.Service
}

// NewApiService creates an api service
func NewApiService(
	cardanoService cardano.Service,
	novelliaDatabaseService novellia_database.Service,
) nvla.DefaultApiServicer {
	return &ApiService {
		cardanoService: cardanoService,
		novelliaDatabaseService: novelliaDatabaseService,
	}
}

// Gets list of products
func (s *ApiService) GetProducts(ctx context.Context, marketId string, organizationId string) (nvla.ImplResponse, error) {
	products, err := s.novelliaDatabaseService.QueryProductIDs(ctx, organizationId, marketId)
	if err != nil {
		err = fmt.Errorf("get products failed at product ID query: %+v", err)
		// TODO: log this in a proper logging stack
		fmt.Printf("GetProducts error: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	
	productsList := nvla.ProductsList{
		Products: products,
	}
	return nvla.Response(200, productsList), nil
}


// Post for list of products details
func (s *ApiService) PostProducts(ctx context.Context, productsList nvla.ProductsList) (nvla.ImplResponse, error) {
	products, err := s.novelliaDatabaseService.QueryAndAddProduct(ctx, productsList.Products)
	if err != nil {
		err = fmt.Errorf("post products failed at product query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	productIDs := []string{}
	for _, p := range products {
		productIDs = append(productIDs, p.Product.ProductId)
	}

	products, err = s.novelliaDatabaseService.QueryAndAddCommission(ctx, productIDs, products)
	if err != nil {
		err = fmt.Errorf("post products failed at commission query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddAttribution(ctx, productIDs, products)
	if err != nil {
		err = fmt.Errorf("post products failed at attribution query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddRemoteResource(ctx, productIDs, products)
	if err != nil {
		err = fmt.Errorf("post products failed at remote resource query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	products, err = s.novelliaDatabaseService.QueryAndAddProductModified(ctx, productIDs, products)
	if err != nil {
		err = fmt.Errorf("post products failed at product modified query: %+v", err)
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	return nvla.Response(200, products), nil
}

// Availability information about service availability
func (s *ApiService) GetStatus(ctx context.Context) (nvla.ImplResponse, error) {
	errorStrings := []string{}

	var status nvla.Status
	status.Maintenance = false
	// TODO: separate microservice being alive vs. some services being down
	status.Status = "UP"
	
	initialized, syncPercentage, err := s.cardanoService.GetStatus(ctx)
	if err != nil {
		status.Cardano = nvla.StatusCardano{
			Initialized: false,
			SyncPercentage: 0,
		}
		errorStrings = append(errorStrings, fmt.Sprintf("%v", err))
	} else {
		status.Cardano = nvla.StatusCardano{
			Initialized: initialized,
			SyncPercentage: syncPercentage,
		}
	}

	status.Errors = errorStrings
	return nvla.Response(200, status), nil
}

// Cardano chain tip information
func (s *ApiService) GetCardanoTip(ctx context.Context) (nvla.ImplResponse, error) {
	blockNumber, epochNumber, err := s.cardanoService.GetTip(ctx)
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
func (s *ApiService) PostWallet(ctx context.Context, wallet nvla.Wallet) (nvla.ImplResponse, error) {
	tokens, err := s.cardanoService.GetAssets(ctx, wallet)
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	tokens, err = s.cardanoService.Add721Metadata(ctx, tokens)
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	return nvla.Response(200, nvla.TokenList{
		Tokens: tokens,
	}), nil
}

// PostTokens
func (s *ApiService) PostTokens(ctx context.Context, tokenSearch nvla.TokenSearch) (nvla.ImplResponse, error) {
	tokens, err := s.cardanoService.QueryTokens(ctx, tokenSearch)
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}
	tokens, err = s.cardanoService.Add721Metadata(ctx, tokens)
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	return nvla.Response(200, nvla.TokenList{
		Tokens: tokens,
	}),nil
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
