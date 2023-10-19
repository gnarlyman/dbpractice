//go:generate mockgen -source handler.go -destination=../../mock/mock_handler.go -package=mock

package handler

import (
	"github.com/gnarlyman/dbpractice/internal/db/repo"
	"github.com/gnarlyman/dbpractice/swagger"
	"github.com/labstack/echo/v4"
)

type IHandler interface {
	// FindUsers (GET /users)
	FindUsers(ctx echo.Context, params swagger.FindUsersParams) error
	// AddUser (POST /users)
	AddUser(ctx echo.Context) error
	// DeleteUser (DELETE /users/{user_id})
	DeleteUser(ctx echo.Context, userId int32) error
	// FindUserById (GET /users/{user_id})
	FindUserById(ctx echo.Context, userId int32) error
	// PatchUser (PATCH /users/{user_id})
	PatchUser(ctx echo.Context, userId int32) error
	// UpdateUser (PUT /users/{user_id})
	UpdateUser(ctx echo.Context, userId int32) error
}

// Handler represents a handler for manage user accounts
type Handler struct {
	userRepo repo.IUserRepo
}

// NewHandler returns a UserHandler ready for use
func NewHandler(userRepo repo.IUserRepo) IHandler {
	return &Handler{
		userRepo: userRepo,
	}
}

// sendDBPracticeError wraps sending of an error in the swagger.Error format
func sendDBPracticeError(ctx echo.Context, code int, message string) error {
	dbPracticeErr := swagger.Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, dbPracticeErr)
	return err
}
