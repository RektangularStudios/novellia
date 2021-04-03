package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/shurcooL/graphql"

	nvla "github.com/RektangularStudios/novellia-sdk/sdk/server/go/v0"
)

type ApiService struct{
	graphqlClient *graphql.Client
}

// NewApiService creates a default api service
func NewApiService(graphqlClient *graphql.Client) nvla.DefaultApiServicer {
	return &ApiService{
		graphqlClient: graphqlClient,
	}
}

// GetWallet - Your GET endpoint
func (s *ApiService) GetWallet(ctx context.Context, walletAddress string) (nvla.ImplResponse, error) {
	// TODO - update GetWallet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	/*
	var query struct {
		PaymentAddress struct {
			address graphql.String
		}
	}
	*/

	/*
	curl \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "{ cardanoDbMeta { initialized syncPercentage }}"}' \
  http://relay1.rektangularstudios.com:3100/graphql
	*/
	/*
	var query struct {
		CardanoDbMeta struct {
			initialized graphql.Boolean
			//syncPercentage graphql.Float
		}
	}
	err := s.graphqlClient.Query(ctx, &query, nil)
	if err != nil {
		return nvla.Response(500, nil), err
	}

	var tokens []nvla.Token
	tokens = append(tokens,
		nvla.Token{
			PolicyId: fmt.Sprintf("%+v", query.CardanoDbMeta.initialized),
	})
	*/
	var tokens []nvla.Token
	tokens = append(tokens,
		nvla.Token{
			PolicyId:    "0xtNVLA",
			Amount:      25,
			Ticker:      "tNVLA",
			Description: "Test tokens for Novellia",
		},
		nvla.Token{
			PolicyId:    "0xADA",
			Amount:      15,
			Ticker:      "ADA",
			Description: "Cardano's ADA Token",
		})
	

	return nvla.Response(200, tokens), nil

	//TODO: Uncomment the next line to return response nvla.Response(400, {}) or use other options such as http.Ok ...
	//return nvla.Response(400, nil),nil

	//return nvla.Response(http.StatusNotImplemented, nil), errors.New("GetWallet method not implemented")
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
