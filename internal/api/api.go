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

// Availability information about Cardano APIs
func (s *ApiService) GetCardanoStatus(ctx context.Context) (nvla.ImplResponse, error) {
	initialized, syncPercentage, err := s.cardanoGraphQLService.GetStatus(ctx)
	
	if err != nil {
		return nvla.Response(500, fmt.Sprintf("error: %v", err)), nil
	}

	resp := nvla.CardanoStatus{
		Initialized: initialized,
		SyncPercentage: syncPercentage,
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
