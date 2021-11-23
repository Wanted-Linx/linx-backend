package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/wanted-linx/linx-backend/api/domain"
	restErr "github.com/wanted-linx/linx-backend/api/errors"
)

type UserHandler struct {
	UserService domain.UserService
}

func NewUserHandler(userService domain.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	reqSignUp := new(domain.SignUpRequest)
	if err := c.Bind(reqSignUp); err != nil {
		restErr.BadRequestErr = err
		return errors.Wrap(err, "잘못된 회원가입 json body 형식입니다.")
	}

	newUser, err := h.UserService.SignUp(reqSignUp)
	if err != nil {
		return err
	}

	return c.JSON(200, newUser)
}

func (h *UserHandler) Login(c echo.Context) error {
	reqLogin := new(domain.LoginRequest)
	if err := c.Bind(reqLogin); err != nil {
		return errors.Wrap(err, "잘못된 로그인 json body 형식입니다.")
	}

	// TODO: 토큰 발행 로직...
	user, err := h.UserService.Login(reqLogin)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.JSON(200, user)
}

func (H *UserHandler) GetUserByID(c echo.Context) error {

	return c.JSON(200, "")
}
