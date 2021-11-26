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

type CompanyHandler struct {
	companySerivce domain.CompanyService
}

func NewCompanyHandler(companySerivce domain.CompanyService) *CompanyHandler {
	return &CompanyHandler{companySerivce: companySerivce}
}

func (h *CompanyHandler) GetCompany(c echo.Context) error {
	// 인증용 id
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("조회 요청 유저 id: ", userID)

	companyID, err := strconv.Atoi(util.GetParams("company_id", c))
	if err != nil {
		return errors.Wrap(err, "잘못된 query string을 입력했습니다.")
	}

	s, err := h.companySerivce.GetCompanyByID(companyID)
	if err != nil {
		return err
	}

	return c.JSON(200, s)
}

func (h *CompanyHandler) GetAllCompanies(c echo.Context) error {
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("조회 요청 유저 id: ", userID)

	limit, err := strconv.Atoi(util.GetQueryParams("limit", c))
	if err != nil {
		return errors.WithMessage(err, "잘못된 query string으로 입력했습니다.")
	}

	offset, err := strconv.Atoi(util.GetQueryParams("offset", c))
	if err != nil {
		return errors.WithMessage(err, "잘못된 query string으로 입력했습니다.")
	}

	companies, err := h.companySerivce.GetAllCompanies(limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(200, companies)
}

func (h *CompanyHandler) UpdateProfile(c echo.Context) error {
	companyID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	reqUpdate := new(domain.CompanyProfileUpdate)
	if err := c.Bind(&reqUpdate); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	company, err := h.companySerivce.UpdateProfile(companyID, reqUpdate)
	if err != nil {
		return err
	}

	return c.JSON(200, company)
}

func (h *CompanyHandler) UploadProfileImage(c echo.Context) error {
	companyID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	reqImage := &domain.CompanyProfileImage{
		Image: form.File["image"],
	}

	if err := c.Bind(&reqImage); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	fileBytes, err := h.companySerivce.UploadProfileImage(companyID, reqImage)
	if err != nil {
		return err
	}

	return c.Blob(200, http.DetectContentType(fileBytes), fileBytes)
}

func (h *CompanyHandler) GetProfileImage(c echo.Context) error {
	companyID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	fileBytes, err := h.companySerivce.GetProfileImage(companyID)
	if err != nil {
		return err
	}

	return c.Blob(200, http.DetectContentType(fileBytes), fileBytes)
}
