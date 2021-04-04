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
			initialized graphql.Boolean
			syncPercentage graphql.Float
		}
	}

	err := s.graphQLClient.Query(ctx, &query, nil)
	if err != nil {
		return false, 0, err
	}

	return bool(query.CardanoDbMeta.initialized), float64(query.CardanoDbMeta.syncPercentage), nil
}
