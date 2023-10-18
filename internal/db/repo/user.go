//go:generate mockgen -source user.go -destination=../../mock/mock_user.go -package=mock

package repo

import (
	"context"

	"github.com/gnarlyman/dbpractice/internal/db/dtomodel"
	"github.com/gnarlyman/dbpractice/internal/db/sql"
)

type IUser interface {
	CreateUser(ctx context.Context, user *dtomodel.User) (*dtomodel.User, error)
	DeleteUser(ctx context.Context, userID int) error
	GetUser(ctx context.Context, userID int) (*dtomodel.User, error)
	GetUserWithPassword(ctx context.Context, userID int) (*dtomodel.User, error)
	ListUsers(ctx context.Context) ([]*dtomodel.User, error)
	UpdateUser(ctx context.Context, user *dtomodel.User) (*dtomodel.User, error)
	PatchUser(ctx context.Context, userID int, userUpdate *dtomodel.User) (*dtomodel.User, error)
}

type UserRepo struct {
	db *sql.Queries
}

func NewUserRepo(db *sql.Queries) IUser {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *dtomodel.User) (*dtomodel.User, error) {
	createUserParams := sql.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	userRow, err := ur.db.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}
	return &dtomodel.User{
		UserID:    int(userRow.UserID),
		Username:  userRow.Username,
		Email:     userRow.Email,
		CreatedAt: userRow.CreatedAt.Time,
		UpdatedAt: userRow.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, userID int) error {
	if err := ur.db.DeleteUser(ctx, int32(userID)); err != nil {
		return err
	}
	return nil
}

// GetUser retrieve a user from the database using the userID
func (ur *UserRepo) GetUser(ctx context.Context, userID int) (*dtomodel.User, error) {
	userRow, err := ur.db.GetUser(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	return &dtomodel.User{
		UserID:    int(userRow.UserID),
		Username:  userRow.Username,
		Email:     userRow.Email,
		CreatedAt: userRow.CreatedAt.Time,
		UpdatedAt: userRow.UpdatedAt.Time,
	}, nil
}

// GetUserWithPassword retrieve a user from the database using the userID, includes password
func (ur *UserRepo) GetUserWithPassword(ctx context.Context, userID int) (*dtomodel.User, error) {
	userRow, err := ur.db.GetUserWithPassword(ctx, int32(userID))
	if err != nil {
		return nil, err
	}
	return &dtomodel.User{
		UserID:    int(userRow.UserID),
		Username:  userRow.Username,
		Email:     userRow.Email,
		Password:  userRow.Password,
		CreatedAt: userRow.CreatedAt.Time,
		UpdatedAt: userRow.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepo) ListUsers(ctx context.Context) ([]*dtomodel.User, error) {
	listUsersRow, err := ur.db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []*dtomodel.User
	for _, userRow := range listUsersRow {
		users = append(users, &dtomodel.User{
			UserID:    int(userRow.UserID),
			Username:  userRow.Username,
			Email:     userRow.Email,
			CreatedAt: userRow.CreatedAt.Time,
			UpdatedAt: userRow.UpdatedAt.Time,
		})
	}

	return users, nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, userUpdate *dtomodel.User) (*dtomodel.User, error) {
	params := sql.UpdateUserParams{
		UserID:   int32(userUpdate.UserID),
		Username: userUpdate.Username,
		Email:    userUpdate.Email,
		Password: userUpdate.Password,
	}

	user, err := ur.db.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &dtomodel.User{
		UserID:    int(user.UserID),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepo) PatchUser(ctx context.Context, userID int, userUpdate *dtomodel.User) (*dtomodel.User, error) {
	user, err := ur.GetUserWithPassword(ctx, userID)
	if err != nil {
		return nil, err
	}

	userUpdate.UserID = userID

	if userUpdate.Email == "" {
		userUpdate.Email = user.Email
	}
	if userUpdate.Username == "" {
		userUpdate.Username = user.Username
	}
	if userUpdate.Password == "" {
		userUpdate.Password = user.Password
	}

	updatedUser, err := ur.UpdateUser(ctx, userUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
