//go:generate mockgen -source user.go -destination=../../mock/mock_user_repo.go -package=mock

package repo

import (
	"context"
	"time"

	"github.com/gnarlyman/dbpractice/internal/db/sql"
	"github.com/gnarlyman/dbpractice/swagger"
)

type IUserRepo interface {
	CreateUser(ctx context.Context, user *swagger.User) (*swagger.User, error)
	DeleteUser(ctx context.Context, userID int32) error
	GetUser(ctx context.Context, userID int32) (*swagger.User, error)
	GetUserWithPassword(ctx context.Context, userID int32) (*swagger.User, error)
	FindUsers(ctx context.Context) ([]*swagger.User, error)
	UpdateUser(ctx context.Context, user *swagger.User) (*swagger.User, error)
	PatchUser(ctx context.Context, userID int32, userUpdate *swagger.User) (*swagger.User, error)
}

type UserRepo struct {
	db *sql.Queries
}

func NewUserRepo(db *sql.Queries) IUserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *swagger.User) (*swagger.User, error) {
	createUserParams := sql.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: *user.Password,
	}
	userRow, err := ur.db.CreateUser(ctx, createUserParams)
	if err != nil {
		return nil, err
	}

	createdAt := userRow.CreatedAt.Time.Format(time.RFC3339)
	updatedAt := userRow.UpdatedAt.Time.Format(time.RFC3339)

	return &swagger.User{
		UserId:    userRow.UserID,
		Username:  userRow.Username,
		Email:     userRow.Email,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}

func (ur *UserRepo) DeleteUser(ctx context.Context, userID int32) error {
	if err := ur.db.DeleteUser(ctx, userID); err != nil {
		return err
	}
	return nil
}

// GetUser retrieve a user from the database using the userID
func (ur *UserRepo) GetUser(ctx context.Context, userID int32) (*swagger.User, error) {
	userRow, err := ur.db.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	createdAt := userRow.CreatedAt.Time.Format(time.RFC3339)
	updatedAt := userRow.UpdatedAt.Time.Format(time.RFC3339)

	return &swagger.User{
		UserId:    userRow.UserID,
		Username:  userRow.Username,
		Email:     userRow.Email,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}

// GetUserWithPassword retrieve a user from the database using the userID, includes password
func (ur *UserRepo) GetUserWithPassword(ctx context.Context, userID int32) (*swagger.User, error) {
	userRow, err := ur.db.GetUserWithPassword(ctx, userID)
	if err != nil {
		return nil, err
	}

	createdAt := userRow.CreatedAt.Time.Format(time.RFC3339)
	updatedAt := userRow.UpdatedAt.Time.Format(time.RFC3339)

	return &swagger.User{
		UserId:    userRow.UserID,
		Username:  userRow.Username,
		Email:     userRow.Email,
		Password:  &userRow.Password,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}

func (ur *UserRepo) FindUsers(ctx context.Context) ([]*swagger.User, error) {
	listUsersRow, err := ur.db.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var users []*swagger.User
	for _, userRow := range listUsersRow {

		createdAt := userRow.CreatedAt.Time.Format(time.RFC3339)
		updatedAt := userRow.UpdatedAt.Time.Format(time.RFC3339)

		users = append(users, &swagger.User{
			UserId:    userRow.UserID,
			Username:  userRow.Username,
			Email:     userRow.Email,
			CreatedAt: &createdAt,
			UpdatedAt: &updatedAt,
		})
	}

	return users, nil
}

func (ur *UserRepo) UpdateUser(ctx context.Context, userUpdate *swagger.User) (*swagger.User, error) {
	params := sql.UpdateUserParams{
		UserID:   userUpdate.UserId,
		Username: userUpdate.Username,
		Email:    userUpdate.Email,
		Password: *userUpdate.Password,
	}

	user, err := ur.db.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	createdAt := user.CreatedAt.Time.Format(time.RFC3339)
	updatedAt := user.UpdatedAt.Time.Format(time.RFC3339)

	return &swagger.User{
		UserId:    user.UserID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
	}, nil
}

func (ur *UserRepo) PatchUser(ctx context.Context, userID int32, userUpdate *swagger.User) (*swagger.User, error) {
	user, err := ur.GetUserWithPassword(ctx, userID)
	if err != nil {
		return nil, err
	}

	userUpdate.UserId = userID

	if userUpdate.Email == "" {
		userUpdate.Email = user.Email
	}
	if userUpdate.Username == "" {
		userUpdate.Username = user.Username
	}
	if userUpdate.Password == nil {
		userUpdate.Password = user.Password
	}

	updatedUser, err := ur.UpdateUser(ctx, userUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
