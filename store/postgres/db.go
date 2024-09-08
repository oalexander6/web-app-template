package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oalexander6/web-app-template/config"
	"github.com/oalexander6/web-app-template/logger"
)

type PostgresStore struct {
	DB *pgxpool.Pool
}

var schema = `
CREATE TABLE IF NOT EXISTS notes (
	id         BIGSERIAL PRIMARY KEY,
	name       TEXT NOT NULL,
	value      TEXT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	deleted    BOOLEAN NOT NULL
);
`

func New(opts config.PostgresConfig) *PostgresStore {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := pgxpool.New(ctx, opts.URI)
	if err != nil {
		logger.Log.Fatal().Msgf("Unable to create pgx connection pool: %s", err)
	}

	if err = conn.Ping(ctx); err != nil {
		logger.Log.Fatal().Msgf("Failed to ping postgres: %s", err)
	}

	conn.Exec(context.Background(), schema)

	return &PostgresStore{
		DB: conn,
	}
}

func (s PostgresStore) Close() {
	s.DB.Close()
}
