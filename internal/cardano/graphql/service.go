package cardano_graphql

import (
	"context"
	"github.com/shurcooL/graphql"
)

type ServiceImpl struct {
	graphQLClient *graphql.Client
}

func New(graphQLClient *graphql.Client) (*ServiceImpl) {
	return &ServiceImpl {
		graphQLClient: graphQLClient,
	}
}

// Returns the initialization status of the remote GraphQL instance, and the sync percentage
func (s *ServiceImpl) Initialized(ctx context.Context) (bool, float64, error) {
	var query struct {
		CardanoDbMeta struct {
			Initialized graphql.Boolean
			SyncPercentage graphql.Float
		}
	}

	err := s.graphQLClient.Query(ctx, &query, nil)
	if err != nil {
		return false, 0, err
	}

	// TODO: verify if casting like this is best practice
	return bool(query.CardanoDbMeta.Initialized), float64(query.CardanoDbMeta.SyncPercentage), nil
}

func (s *ServiceImpl) GetAssets(ctx context.Context, paymentAddress string) (error) {
	// TODO: fill stub
	return nil
}
