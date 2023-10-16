package dbpractice

import (
	"context"
	"log"

	"github.com/gnarlyman/dbpractice/internal/db"
	"github.com/gnarlyman/dbpractice/internal/db/repository"
	"github.com/gnarlyman/dbpractice/internal/http"
	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/vingarcia/ksql"
)

type DbPractice struct {
	conn       *ksql.DB
	httpServer *http.Server
}

func NewDbPractice() *DbPractice {
	cfg := NewConfig()

	conn, err := db.NewDB(cfg.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(conn)
	userHandler := handler.NewUserHandler(userRepository)
	httpServer := http.NewHttpServer(cfg.ListenAddr, userHandler)

	return &DbPractice{
		conn:       conn,
		httpServer: httpServer,
	}
}

func (d *DbPractice) Start(ctx context.Context) {
	go func() {
		d.httpServer.StartServer()
	}()
}

func (d *DbPractice) Stop() {
	d.conn.Close()
}
