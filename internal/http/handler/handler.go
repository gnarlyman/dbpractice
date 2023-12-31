//go:generate mockgen -source handler.go -destination=../../mock/mock_handler.go -package=mock

package handler

import (
	"net/http"

	"github.com/gnarlyman/dbpractice/internal/db/repo"
	"github.com/gnarlyman/dbpractice/pkg/swagger"
	"github.com/labstack/echo/v4"
)

type IHandler interface {
	// GetApiV1SwaggerJson (GET /api/v1/swagger.json)
	GetApiV1SwaggerJson(ctx echo.Context) error
	// HeadApiV1Users (HEAD /api/v1/users)
	HeadApiV1Users(ctx echo.Context) error
	// OptionsApiV1Users (OPTIONS /api/v1/users)
	OptionsApiV1Users(ctx echo.Context) error
	// HeadApiV1UsersUserId (HEAD /api/v1/users/{user_id})
	HeadApiV1UsersUserId(ctx echo.Context, userId int32) error
	// OptionsApiV1UsersUserId (OPTIONS /api/v1/users/{user_id})
	OptionsApiV1UsersUserId(ctx echo.Context, userId int32) error
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

// GetApiV1SwaggerJson server the swagger as json
func (h *Handler) GetApiV1SwaggerJson(ctx echo.Context) error {
	sw, err := swagger.GetSwagger()
	if err != nil {
		return sendDBPracticeError(ctx, http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, sw)
}
