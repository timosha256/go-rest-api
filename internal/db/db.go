package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTables(ctx context.Context, conn *pgx.Conn) (pgx.Rows, error) {
	return conn.Query(
		ctx,
		`
            CREATE TABLE IF NOT EXISTS users (
                id SERIAL PRIMARY KEY,
                firstname VARCHAR(100) NOT NULL,
                lastname VARCHAR(100) NOT NULL,
                age INT,
                email VARCHAR(100) NOT NULL UNIQUE,
                country VARCHAR(100) NOT NULL,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            )
        `,
	)
}
