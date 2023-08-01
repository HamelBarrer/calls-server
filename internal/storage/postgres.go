package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type Postgres struct {
	db *pgxpool.Pool
}

func NewPostgres() (*Postgres, error) {
	viper.SetConfigName(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	username := viper.Get("DATABASE_USER")
	password := viper.Get("DATABASE_PASSWORD")
	host := viper.Get("DATABASE_HOST")
	port := viper.Get("DATABASE_PORT")
	db := viper.Get("DATABASE_DB")

	urlconn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, db)

	p, err := pgxpool.New(context.Background(), urlconn)
	if err != nil {
		return nil, err
	}

	return &Postgres{p}, nil
}
