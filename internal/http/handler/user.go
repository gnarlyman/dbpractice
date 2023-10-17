package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gnarlyman/dbpractice/internal/db/dbmodel"
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

// AllUser  returns all users to client
func (h *UserHandler) AllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepository.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UserById returns a given user to client
func (h *UserHandler) UserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.userRepository.GetUserById(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// NewUser creates a new user from post body
func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var user dbmodel.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser, err := h.userRepository.NewUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

// UpdateUser updates a user with new settings
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user dbmodel.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepository.UpdateUser(r.Context(), userId, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

// DeleteUser removes user based on their username
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	if err := h.userRepository.DeleteUser(r.Context(), userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
