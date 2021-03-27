package api_service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	novellia_api "github.com/RektangularStudios/novellia/generated/novellia-api"
)

type ApiService struct{}

// NewApiService creates a default api service
func NewApiService() novellia_api.DefaultApiServicer {
	return &ApiService{}
}

// GetWallet - Your GET endpoint
func (s *ApiService) GetWallet(ctx context.Context, getWalletRequest novellia_api.GetWalletRequest) (novellia_api.ImplResponse, error) {
	fmt.Printf("GET WALLET")

	// TODO - update GetWallet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Token{}) or use other options such as http.Ok ...
	//return Response(200, []Token{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("GetWallet method not implemented")
}

// GetWorkflowMinterNvla -
func (s *ApiService) GetWorkflowMinterNvla(ctx context.Context) (novellia_api.ImplResponse, error) {
	// TODO - update GetWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, WorkflowInformation{}) or use other options such as http.Ok ...
	//return Response(200, WorkflowInformation{}), nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("GetWorkflowMinterNvla method not implemented")
}

// PostWallet -
func (s *ApiService) PostWallet(ctx context.Context, cardanoTransaction novellia_api.CardanoTransaction) (novellia_api.ImplResponse, error) {
	// TODO - update PostWallet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("PostWallet method not implemented")
}

// PostWorkflowMinterNvla -
func (s *ApiService) PostWorkflowMinterNvla(ctx context.Context, minterRequest novellia_api.MinterRequest) (novellia_api.ImplResponse, error) {
	// TODO - update PostWorkflowMinterNvla with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, CardanoTransaction{}) or use other options such as http.Ok ...
	//return Response(200, CardanoTransaction{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return novellia_api.Response(http.StatusNotImplemented, nil), errors.New("PostWorkflowMinterNvla method not implemented")
}
