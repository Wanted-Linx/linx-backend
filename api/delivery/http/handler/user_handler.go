package handler

import (
	"log"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	restErr "github.com/Wanted-Linx/linx-backend/api/errors"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type UserHandler struct {
	UserService    domain.UserService
	StudentSerivce domain.StudentService
}

func NewUserHandler(userService domain.UserService, studentService domain.StudentService) *UserHandler {
	return &UserHandler{
		UserService:    userService,
		StudentSerivce: studentService,
	}
}

func (h *UserHandler) SignUp(c echo.Context) error {
	reqSignUp := new(domain.SignUpRequest)
	if err := c.Bind(reqSignUp); err != nil {
		restErr.BadRequestErr = err
		return errors.Wrap(err, "잘못된 회원가입 json body 입니다.")
	}

	newUser, err := h.UserService.SignUp(reqSignUp)
	if err != nil {
		return err
	}

	if newUser.Kind == "student" {
		newStudent, err := h.StudentSerivce.Save(newUser.ID, reqSignUp)
		if err != nil {
			return err
		}
		return c.JSON(200, newStudent)
	}

	return c.JSON(200, "")

	// newCompany, err := h.CompanyService.Save(newUser.ID, reqSignUp)
	// if err != nil {
	// 	return nil
	// }
	// return c.JSON(200, newCompany)
}

func (h *UserHandler) Login(c echo.Context) error {
	userID, _ := util.GetRequestUserID(c)
	log.Println(userID)
	reqLogin := new(domain.LoginRequest)
	if err := c.Bind(reqLogin); err != nil {
		return errors.Wrap(err, "잘못된 로그인 json body 입니다.")
	}

	// TODO: 토큰 발행 로직...
	user, err := h.UserService.Login(reqLogin)
	if err != nil {
		return err
	}

	return c.JSON(200, user)
}

func (H *UserHandler) GetUserByID(c echo.Context) error {

	return c.JSON(200, "")
}
