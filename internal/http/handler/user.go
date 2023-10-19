package handler

import (
	"errors"
	"net/http"

	"github.com/gnarlyman/dbpractice/swagger"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

// FindUsers returns any swagger.User objects to client
func (h *Handler) FindUsers(ctx echo.Context, params swagger.FindUsersParams) error {
	users, err := h.userRepo.FindUsers(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, users)
}

// FindUserById returns a swagger.User to client from a given userId
func (h *Handler) FindUserById(ctx echo.Context, userId int32) error {
	user, err := h.userRepo.GetUser(ctx.Request().Context(), userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sendDBPracticeError(ctx, http.StatusNotFound, "user not found")
		}
		return err
	}

	return ctx.JSON(http.StatusOK, user)
}

// AddUser creates a new swagger.User from post body
func (h *Handler) AddUser(ctx echo.Context) error {
	var newUser swagger.NewUser
	err := ctx.Bind(&newUser)
	if err != nil {
		return sendDBPracticeError(ctx, http.StatusBadRequest, "Invalid format for NewUser")
	}

	var user swagger.User
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = newUser.Password

	createdUser, err := h.userRepo.CreateUser(ctx.Request().Context(), &user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, createdUser)
}

// UpdateUser updates a swagger.User with new settings
func (h *Handler) UpdateUser(ctx echo.Context, userId int32) error {
	var newUser swagger.NewUser
	err := ctx.Bind(&newUser)
	if err != nil {
		return sendDBPracticeError(ctx, http.StatusBadRequest, "Invalid format for NewUser")
	}

	var user swagger.User
	user.UserId = userId
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = newUser.Password

	updatedUser, err := h.userRepo.UpdateUser(ctx.Request().Context(), &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sendDBPracticeError(ctx, http.StatusNotFound, "user not found")
		}
		return err
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

// PatchUser updates only specified fields of swagger.User
func (h *Handler) PatchUser(ctx echo.Context, userId int32) error {
	var newUser swagger.NewUser
	err := ctx.Bind(&newUser)
	if err != nil {
		return sendDBPracticeError(ctx, http.StatusBadRequest, "Invalid format for NewUser")
	}

	var user swagger.User
	user.UserId = userId
	user.Username = newUser.Username
	user.Email = newUser.Email
	user.Password = newUser.Password

	updatedUser, err := h.userRepo.PatchUser(ctx.Request().Context(), userId, &user)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sendDBPracticeError(ctx, http.StatusNotFound, "user not found")
		}
		return err
	}

	return ctx.JSON(http.StatusOK, updatedUser)
}

// DeleteUser removes swagger.User based on their username
func (h *Handler) DeleteUser(ctx echo.Context, userId int32) error {
	if err := h.userRepo.DeleteUser(ctx.Request().Context(), userId); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return sendDBPracticeError(ctx, http.StatusNotFound, "user not found")
		}
		return err
	}
	return nil
}
