package dbpractice

import (
	"context"
	"log"

	"github.com/gnarlyman/dbpractice/internal/db"
	"github.com/gnarlyman/dbpractice/internal/http"
	"github.com/gnarlyman/dbpractice/internal/http/handler"
)

type DbPractice struct {
	db         db.IDB
	httpServer *http.Server
}

func NewDbPractice() *DbPractice {
	cfg := NewConfig()

	appDb, err := db.NewDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := handler.NewUserHandler(appDb)
	httpServer := http.NewHttpServer(cfg.ListenAddr, userHandler)

	return &DbPractice{
		db:         appDb,
		httpServer: httpServer,
	}
}

func (d *DbPractice) Start(ctx context.Context) {
	go func() {
		d.httpServer.StartServer()
	}()
}

func (d *DbPractice) Stop(ctx context.Context) {
	d.db.Stop(ctx)
}
