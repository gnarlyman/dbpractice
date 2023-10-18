//go:generate mockgen -source db.go -destination=../mock/mock_db.go -package=mock

package db

import (
	"context"
	"log"

	"github.com/gnarlyman/dbpractice/internal/db/repo"
	"github.com/gnarlyman/dbpractice/internal/db/sql"
	"github.com/jackc/pgx/v5"
)

type IDB interface {
	GetUserRepo() repo.IUser
	Stop(ctx context.Context) error
}

type DB struct {
	conn     *pgx.Conn
	db       *sql.Queries
	UserRepo repo.IUser
}

func NewDB(databaseUrl string) (IDB, error) {
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	db := sql.New(conn)

	app := &DB{
		conn:     conn,
		UserRepo: repo.NewUserRepo(db),
	}

	return app, nil
}

func (a *DB) GetUserRepo() repo.IUser {
	return a.UserRepo
}

func (a *DB) Stop(ctx context.Context) error {
	return a.conn.Close(ctx)
}
