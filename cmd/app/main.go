package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/timosha256/go-rest-api/internal/db"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env")
	}

	conn, err := pgx.Connect(ctx, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("Could not connect to DB", err)
	}

	// pool, err := pgxpool.New(ctx, os.Getenv("DB_URI"))
	// if err != nil {
	// 	log.Fatal("Could not create DB pool")
	// }

	_, err = db.CreateTables(ctx, conn)
	if err != nil {
		log.Fatal("Could not create tables", err)
	}
}
