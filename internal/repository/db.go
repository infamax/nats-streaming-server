package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
}

func NewDB(ctx context.Context, connectionString string) (*db, error) {
	pool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return &db{
		pool: pool,
	}, nil
}
