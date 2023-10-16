package http

import (
	"net/http"

	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/gorilla/mux"
)

type Server struct {
	userHandler *handler.UserHandler
	listenAddr  string
}

func NewHttpServer(listenAddr string, userHandler *handler.UserHandler) *Server {
	return &Server{
		userHandler: userHandler,
		listenAddr:  listenAddr,
	}
}

func (s *Server) StartServer() {
	r := mux.NewRouter()
	r.HandleFunc(
		"/user", s.userHandler.AllUserHandler).
		Methods(http.MethodGet)
	r.HandleFunc(
		"/user/{username}", s.userHandler.UserByUsernameHandler).
		Methods(http.MethodGet)
	http.ListenAndServe(s.listenAddr, r)
}
