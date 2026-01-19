package infrastructure

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(dataSource string) (*sql.DB, error) {
	return sql.Open("postgres", dataSource)
}