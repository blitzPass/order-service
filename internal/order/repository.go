package order

import "database/sql"

type Repository interface {
	Create() (*Order, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create()(*Order, error) {
	return nil, nil;
}