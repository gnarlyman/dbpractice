package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gnarlyman/dbpractice/internal/db/dbmodels"
	"github.com/vingarcia/ksql"
)

// UserRepository represents a user repository for manage database users
type UserRepository struct {
	db *ksql.DB
}

// NewUserRepository returns a UserRepository ready for use
func NewUserRepository(db *ksql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// GetUserByUsername retrieves a user by username
func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*dbmodels.User, error) {
	var user dbmodels.User
	query := "SELECT user_id, username, email, created_at FROM users WHERE username=$1"
	if err := r.db.QueryOne(ctx, &user, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*dbmodels.User, error) {
	var users []*dbmodels.User
	query := "SELECT user_id, username, email, created_at FROM users"
	if err := r.db.Query(ctx, &users, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}
