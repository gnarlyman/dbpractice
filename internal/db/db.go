package db

import (
	"context"
	"log"

	"github.com/gnarlyman/dbpractice/internal/db/dtomodel"
	"github.com/gnarlyman/dbpractice/internal/db/sql"
	"github.com/jackc/pgx/v5"
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

type IDB interface {
	IUser
	Stop(ctx context.Context) error
}

func NewDB(databaseUrl string) (IDB, error) {
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	app := &App{
		conn: conn,
		db:   sql.New(conn),
	}

	return app, nil
}

func (a App) Stop(ctx context.Context) error {
	return a.conn.Close(ctx)
}

type App struct {
	conn *pgx.Conn
	db   *sql.Queries
}

func (a App) CreateUser(ctx context.Context, user *dtomodel.User) (*dtomodel.User, error) {
	createUserParams := sql.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	userRow, err := a.db.CreateUser(ctx, createUserParams)
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

func (a App) DeleteUser(ctx context.Context, userID int) error {
	if err := a.db.DeleteUser(ctx, int32(userID)); err != nil {
		return err
	}
	return nil
}

// GetUser retrieve a user from the database using the userID
func (a App) GetUser(ctx context.Context, userID int) (*dtomodel.User, error) {
	userRow, err := a.db.GetUser(ctx, int32(userID))
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
func (a App) GetUserWithPassword(ctx context.Context, userID int) (*dtomodel.User, error) {
	userRow, err := a.db.GetUserWithPassword(ctx, int32(userID))
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

func (a App) ListUsers(ctx context.Context) ([]*dtomodel.User, error) {
	listUsersRow, err := a.db.ListUsers(ctx)
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

func (a App) UpdateUser(ctx context.Context, user *dtomodel.User) (*dtomodel.User, error) {
	params := sql.UpdateUserParams{
		UserID:   int32(user.UserID),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	u, err := a.db.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}

	return &dtomodel.User{
		UserID:    int(u.UserID),
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}, nil
}

func (a App) PatchUser(ctx context.Context, userID int, userUpdate *dtomodel.User) (*dtomodel.User, error) {
	user, err := a.GetUserWithPassword(ctx, userID)
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

	updatedUser, err := a.UpdateUser(ctx, userUpdate)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
