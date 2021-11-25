package handler

import (
	"net/http"
	"strconv"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type StudentHandler struct {
	studentService domain.StudentService
}

func NewStudentHandler(studentService domain.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

func (h *StudentHandler) GetStudent(c echo.Context) error {
	// 인증용 id
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("조회 요청 유저 id: ", userID)

	ownerID, err := strconv.Atoi(util.GetQueryParams("owner", c))
	if err != nil {
		return errors.Wrap(err, "잘못된 query string을 입력했습니다.")
	}

	s, err := h.studentService.GetStudentByID(ownerID)
	if err != nil {
		return err
	}

	return c.JSON(200, s)
}

func (h *StudentHandler) UpdateProfile(c echo.Context) error {
	studentID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	reqUpdate := new(domain.StudentProfileUpdate)
	if err := c.Bind(&reqUpdate); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	s, err := h.studentService.UpdateProfile(studentID, reqUpdate)
	if err != nil {
		return err
	}

	return c.JSON(200, s)
}

func (h *StudentHandler) UploadProfileImage(c echo.Context) error {
	studentID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	reqImage := &domain.StudentProfileImage{
		Image: form.File["image"],
	}

	if err := c.Bind(&reqImage); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	fileBytes, err := h.studentService.UploadProfileImage(studentID, reqImage)
	if err != nil {
		return err
	}

	return c.Blob(200, "image/png", fileBytes)
}

func (h *StudentHandler) GetProfileImage(c echo.Context) error {
	studentID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}
	fileBytes, err := h.studentService.GetProfileImage(studentID)
	if err != nil {
		return err
	}

	return c.Blob(200, http.DetectContentType(fileBytes), fileBytes)
}
