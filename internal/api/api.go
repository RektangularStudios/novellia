package api

import (
	"context"
	"errors"
	"net/http"

	novellia_api "github.com/RektangularStudios/novellia/generated/novellia-api"
)

type ApiService struct{}

// NewApiService creates a default api service
func NewApiService() novellia_api.DefaultApiServicer {
	return &ApiService{}
}

// GetWallet - Your GET endpoint
func (s *ApiService) GetWallet(ctx context.Context, walletAddress string) (novellia_api.ImplResponse, error) {
	// TODO - update GetWallet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	var tokens []novellia_api.Token
	tokens = append(tokens,
		novellia_api.Token{
			PolicyId:    "0xtNVLA",
			Amount:      25,
			Ticker:      "tNVLA",
			Description: "Test tokens for Novellia",
		},
		novellia_api.Token{
			PolicyId:    "0xADA",
			Amount:      15,
			Ticker:      "ADA",
			Description: "Cardano's ADA Token",
		})
	return novellia_api.Response(200, tokens), nil

	//TODO: Uncomment the next line to return response novellia_api.Response(400, {}) or use other options such as http.Ok ...
	//return novellia_api.Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("GetWallet method not implemented")
}

// GetWorkflowMinterNvla -
func (s *ApiService) GetWorkflowMinterNvla(ctx context.Context) (novellia_api.ImplResponse, error) {
	// TODO - update GetWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response novellia_api.Response(200, WorkflowInformation{}) or use other options such as http.Ok ...
	//return novellia_api.Response(200, WorkflowInformation{}), nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("GetWorkflowMinterNvla method not implemented")
}

// PostCardanoTransaction -
func (s *ApiService) PostCardanoTransaction(ctx context.Context, cardanoTransaction novellia_api.CardanoTransaction) (novellia_api.ImplResponse, error) {
	// TODO - update PostCardanoTransaction with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response novellia_api.Response(200, {}) or use other options such as http.Ok ...
	//return novellia_api.Response(200, nil),nil

	//TODO: Uncomment the next line to return response novellia_api.Response(400, {}) or use other options such as http.Ok ...
	//return novellia_api.Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("PostCardanoTransaction method not implemented")
}

// PostWorkflowMinterNvla -
func (s *ApiService) PostWorkflowMinterNvla(ctx context.Context, minterInfo novellia_api.MinterInfo) (novellia_api.ImplResponse, error) {
	// TODO - update PostWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response novellia_api.Response(200, CardanoTransaction{}) or use other options such as http.Ok ...
	//return novellia_api.Response(200, CardanoTransaction{}), nil

	//TODO: Uncomment the next line to return response novellia_api.Response(400, {}) or use other options such as http.Ok ...
	//return novellia_api.Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("PostWorkflowMinterNvla method not implemented")
}
