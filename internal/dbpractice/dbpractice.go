package dbpractice

import (
	"context"
	"log"

	"github.com/gnarlyman/dbpractice/internal/db"
	"github.com/gnarlyman/dbpractice/internal/http"
	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/sirupsen/logrus"
)

type DbPractice struct {
	db         db.IDB
	httpServer *http.Server
}

func NewDbPractice() *DbPractice {
	cfg := NewConfig()
	lgr := logrus.New()
	lgr.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	appDb, err := db.NewDB(cfg.DatabaseUrl, lgr)
	if err != nil {
		log.Fatal(err)
	}

	apiHandler := handler.NewHandler(appDb.GetUserRepo())
	httpServer := http.NewHttpServer(cfg.ListenAddr, apiHandler)

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

func (d *DbPractice) Stop() {
	d.db.Stop()
}
