package db

import "github.com/jackc/pgx/v5/pgxpool"

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new Store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(connPool),
		connPool: connPool,
	}
}
