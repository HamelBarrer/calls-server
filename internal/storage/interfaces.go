package storage

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Storage interface {
	Query(string, ...interface{}) (pgx.Rows, error)
	QueryRow(string, ...interface{}) pgx.Row
	Exec(query string, args ...interface{}) (pgconn.CommandTag, error)
}
