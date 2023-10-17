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
		"/user", s.userHandler.AllUser).
		Methods(http.MethodGet)
	r.HandleFunc(
		"/user/{user_id}", s.userHandler.UserById).
		Methods(http.MethodGet)
	r.HandleFunc(
		"/user", s.userHandler.NewUser).
		Methods(http.MethodPost)
	r.HandleFunc(
		"/user", s.userHandler.UpdateUser).
		Methods(http.MethodPut)
	r.HandleFunc(
		"/user/{user_id}", s.userHandler.PatchUser).
		Methods(http.MethodPatch)
	r.HandleFunc(
		"/user/{user_id}", s.userHandler.DeleteUser).
		Methods(http.MethodDelete)

	http.ListenAndServe(s.listenAddr, r)
}
