package http

import (
	"fmt"
	"log"
	"os"

	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/gnarlyman/dbpractice/pkg/swagger"
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

func CORS() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) (returnErr error) {
			ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Response().Header().Set("Access-Control-Allow-Methods", "HEAD, OPTIONS, POST, GET, PUT, PATCH, DELETE")
			ctx.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
			return next(ctx)
		}
	}
}

func (s *Server) StartServer() {
	sw, err := swagger.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	fmt.Println(sw.Paths)
	sw.Servers = nil

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(oapimiddleware.OapiRequestValidator(sw))
	e.Use(CORS())

	swagger.RegisterHandlers(e, s.handler)
	log.Fatal(e.Start(s.listenAddr))
}
