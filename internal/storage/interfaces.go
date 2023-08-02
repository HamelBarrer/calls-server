package storage

import "github.com/jackc/pgx/v5"

type Storage interface {
	Query(string, ...interface{}) (pgx.Rows, error)
	QueryRow(string, ...interface{}) pgx.Row
}
