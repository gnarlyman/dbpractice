package http

import (
	"log"

	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/gnarlyman/dbpractice/swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	handler    handler.IHandler
	listenAddr string
}

func NewHttpServer(listenAddr string, handler handler.IHandler) *Server {
	return &Server{
		handler:    handler,
		listenAddr: listenAddr,
	}
}

func (s *Server) StartServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	v1 := e.Group("/v1")

	swagger.RegisterHandlers(v1, s.handler)
	log.Fatal(e.Start(s.listenAddr))
}
