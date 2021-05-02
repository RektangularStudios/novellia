package novellia_database

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v4"
)

type ServiceImpl struct {
	conn *pgx.Conn
}

func New(ctx context.Context, username, password, host, database_name string) (*ServiceImpl, error) {
	// url like "postgresql://username:password@localhost:5432/database_name"
	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s/%s", username, password, host, database_name)
	conn, err := pgx.Connect(ctx, databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Postgres: %v", err)
	}

	return &ServiceImpl {
		conn: conn, 
	}, nil
}

func (s *ServiceImpl) ListProductAttribution(ctx context.Context) error {
	query := "SELECT * from product_attribution"
	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return err
	}

	fmt.Printf("list: %+v", rows)
	return nil
}

func (s *ServiceImpl) Close(ctx context.Context) {
	s.conn.Close(ctx)
}
