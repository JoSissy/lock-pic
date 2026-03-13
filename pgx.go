package main

import (
	"context"

	"github.com/google/uuid"
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

func SaveImage(ctx context.Context, id uuid.UUID, filename, mimeType string, sizeBytes int64, data []byte, availableAt string) (string, error) {
	query := `
		INSERT INTO images (id, filename, mime_type, size_bytes, data, available_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var idStr string

	err := db.pool.QueryRow(ctx, query, id, filename, mimeType, sizeBytes, data, availableAt).Scan(&idStr)
	if err != nil {
		logrus.Error("Failed to save image: ", err)
		return "", err
	}

	return idStr, nil
}
