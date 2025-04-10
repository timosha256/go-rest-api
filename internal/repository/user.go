package repository

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/timosha256/go-rest-api/internal/models"
)

type UserRepository struct {
	DB   *pgx.Conn
	Pool *pgxpool.Pool
}

func NewUserRepository(db *pgx.Conn, pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		DB:   db,
		Pool: pool,
	}
}

func (r *UserRepository) Create(ctx context.Context, user models.User) {
	sql := "INSERT INTO users (id, username)"
}
