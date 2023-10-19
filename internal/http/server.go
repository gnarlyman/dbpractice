package http

import (
	"fmt"
	"log"
	"os"

	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/gnarlyman/dbpractice/swagger"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	oapimiddleware "github.com/oapi-codegen/echo-middleware"
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
	sw, err := swagger.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	sw.Servers = nil

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(oapimiddleware.OapiRequestValidator(sw))

	//v1 := e.Group("/v1")

	swagger.RegisterHandlers(e, s.handler)
	log.Fatal(e.Start(s.listenAddr))
}
