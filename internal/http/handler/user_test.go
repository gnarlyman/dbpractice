package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gnarlyman/dbpractice/internal/http/handler"
	"github.com/gnarlyman/dbpractice/internal/mock"
	"github.com/gnarlyman/dbpractice/pkg/swagger"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestFindUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockUserRepo := mock.NewMockIUserRepo(ctrl)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserRepo.EXPECT().FindUsers(c.Request().Context()).
		Return([]*swagger.User{}, nil)

	h := handler.NewHandler(mockUserRepo)

	err := h.FindUsers(c, swagger.FindUsersParams{
		Username: nil,
		Limit:    nil,
	})

	assert.NoError(t, err)
	assert.Equal(t, `[]`, strings.TrimSpace(rec.Body.String()))
}
