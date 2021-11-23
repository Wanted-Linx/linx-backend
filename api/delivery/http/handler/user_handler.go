package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/wanted-linx/linx-backend/api/domain"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	return c.JSON(200, "")
}

func (H *UserHandler) GetUserByID(c echo.Context) error {
	return c.JSON(200, "")
}
