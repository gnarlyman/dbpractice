//go:generate mockgen -source handler.go -destination=../mock/mock_handler.go -package=mock

package handler

import (
	"github.com/gnarlyman/dbpractice/internal/db/repo"
	"github.com/gnarlyman/dbpractice/swagger"
	"github.com/labstack/echo/v4"
)

type IHandler interface {
	FindUsers(ctx echo.Context, params swagger.FindUsersParams) error
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
