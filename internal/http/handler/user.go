package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gnarlyman/dbpractice/internal/db"
	"github.com/gnarlyman/dbpractice/internal/db/dtomodel"
	"github.com/gorilla/mux"
)

// UserHandler represents a handler for manage user accounts
type UserHandler struct {
	repo db.IDB
}

// NewUserHandler returns a UserHandler ready for use
func NewUserHandler(repo db.IDB) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

// AllUser returns all dtomodel.User objects to client
func (h *UserHandler) AllUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.ListUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// UserById returns a given dtomodel.User to client
func (h *UserHandler) UserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUser(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// NewUser creates a new dtomodel.User from post body
func (h *UserHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	var user dtomodel.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdUser, err := h.repo.CreateUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

// UpdateUser updates a dtomodel.User with new settings
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user dtomodel.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdUser, err := h.repo.UpdateUser(r.Context(), &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

// PatchUser updates only specified fields of dtomodel.User
func (h *UserHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var user dtomodel.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedUser, err := h.repo.PatchUser(r.Context(), userId, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser removes dtomodel.User based on their username
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}
	if err := h.repo.DeleteUser(r.Context(), userId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
