//go:generate mockgen -source db.go -destination=../mock/mock_db.go -package=mock

package db

import (
	"context"

	"github.com/gnarlyman/dbpractice/internal/db/repo"
	"github.com/gnarlyman/dbpractice/internal/db/sql"
	"github.com/gnarlyman/dbpractice/pgx5Logger"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type IDB interface {
	GetUserRepo() repo.IUserRepo
	Stop()
}

type DB struct {
	conn     *pgxpool.Pool
	db       *sql.Queries
	UserRepo repo.IUserRepo
}

func NewDB(databaseUrl string, lgr *logrus.Logger) (IDB, error) {
	logger := pgx5Logger.NewLogger(lgr)

	pxgConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	pxgConfig.ConnConfig.Tracer = &tracelog.TraceLog{Logger: logger, LogLevel: tracelog.LogLevelDebug}

	conn, err := pgxpool.NewWithConfig(context.Background(), pxgConfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	db := sql.New(conn)

	app := &DB{
		conn:     conn,
		UserRepo: repo.NewUserRepo(db),
	}

	return app, nil
}

func (a *DB) GetUserRepo() repo.IUserRepo {
	return a.UserRepo
}

func (a *DB) Stop() {
	a.conn.Close()
}
