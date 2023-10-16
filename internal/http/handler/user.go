package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gnarlyman/dbpractice/internal/db/repository"
	"github.com/gorilla/mux"
)

// UserHandler represents a handler for manage user accounts
type UserHandler struct {
	userRepository *repository.UserRepository
}

// NewUserHandler returns a UserHandler ready for use
func NewUserHandler(userRepository *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

// AllUserHandler  returns all users to client
func (h *UserHandler) AllUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UserByUsernameHandler returns a given user to client
func (h *UserHandler) UserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := h.userRepository.GetUserByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
