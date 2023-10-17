package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/gnarlyman/dbpractice/internal/db/dbmodel"
	"github.com/vingarcia/ksql"
)

// UserRepository represents a user repository for manage database users
type UserRepository struct {
	db    *ksql.DB
	table ksql.Table
}

// NewUserRepository returns a UserRepository ready for use
func NewUserRepository(db *ksql.DB) *UserRepository {
	return &UserRepository{
		db:    db,
		table: ksql.NewTable("users", "user_id"),
	}
}

// GetUserById retrieves a user by username
func (r *UserRepository) GetUserById(ctx context.Context, userId int) (*dbmodel.User, error) {
	var user dbmodel.User
	query := `SELECT user_id, username, email, created_at, updated_at FROM users WHERE user_id=$1`
	if err := r.db.QueryOne(ctx, &user, query, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetAllUsers retrieves all users
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]*dbmodel.User, error) {
	var users []*dbmodel.User
	query := `SELECT user_id, username, email, created_at, updated_at FROM users`
	if err := r.db.Query(ctx, &users, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return users, nil
}

// NewUser creates a new user in the database
func (r *UserRepository) NewUser(ctx context.Context, user *dbmodel.User) (*dbmodel.User, error) {
	if err := r.db.Insert(ctx, r.table, user); err != nil {
		return nil, err
	}
	user, err := r.GetUserById(ctx, user.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(ctx context.Context, userId int, user *dbmodel.User) (*dbmodel.User, error) {
	user.UserID = userId
	if err := r.db.Patch(ctx, r.table, user); err != nil {
		return nil, err
	}
	user, err := r.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser removes the user from the database
func (r *UserRepository) DeleteUser(ctx context.Context, userId int) error {
	if err := r.db.Delete(ctx, r.table, userId); err != nil {
		return err
	}
	return nil
}
