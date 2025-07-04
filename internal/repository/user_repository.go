package repository

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ridhogaa/go-jwt-auth/internal/model"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user model.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.db.Exec(context.Background(), query, user.Username, user.Password)
	return err
}

func (r *UserRepository) FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := r.db.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err == pgx.ErrNoRows {
		return model.User{}, sql.ErrNoRows
	}
	return user, err
}
