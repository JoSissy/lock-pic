package main

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

var db *PostgresConn

type PostgresConn struct {
	pool *pgxpool.Pool
}

func PostgresConnect(ctx context.Context) error {
	logrus.Info("Connecting to database...")

	pool, err := pgxpool.New(ctx, c.DatabaseURL)
	if err != nil {
		logrus.Error("Unable to connect to database: ", err)
		return err
	}
	defer pool.Close()

	db = &PostgresConn{pool: pool}

	return nil
}
