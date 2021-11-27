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

type ClubHandler struct {
	clubService domain.ClubService
}

func NewClubHandler(clubService domain.ClubService) *ClubHandler {
	return &ClubHandler{clubService: clubService}
}

func (h *ClubHandler) CreateClub(c echo.Context) error {
	leaderID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	reqClub := new(domain.ClubCreateRequest)
	if err := c.Bind(reqClub); err != nil {
		return errors.Wrap(err, "잘못된 동아리 json body 입니다.")
	}

	newClub, err := h.clubService.CreateClub(leaderID, reqClub)
	if err != nil {
		return err
	}

	return c.JSON(200, newClub)
}

func (h *ClubHandler) GetClubByID(c echo.Context) error {
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청한 유저 id:", userID)
	clubID, err := strconv.Atoi(util.GetParams("club_id", c))
	if err != nil {
		return errors.WithStack(err)
	}

	club, err := h.clubService.GetClubByID(clubID)
	if err != nil {
		return err
	}

	return c.JSON(200, club)
}

func (h *ClubHandler) GetAllClubs(c echo.Context) error {
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청한 유저 id:", userID)
	limit, err := strconv.Atoi(util.GetQueryParams("limit", c))
	if err != nil {
		return errors.WithMessage(err, "잘못된 query string으로 입력했습니다.")
	}

	offset, err := strconv.Atoi(util.GetQueryParams("offset", c))
	if err != nil {
		return errors.WithMessage(err, "잘못된 query string으로 입력했습니다.")
	}

	clubs, err := h.clubService.GetAllClubs(limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(200, clubs)
}

func (h *ClubHandler) UploadProfileImage(c echo.Context) error {
	clubID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	reqImage := &domain.ProfileImageRequest{
		Image: form.File["image"],
	}

	if err := c.Bind(&reqImage); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	fileBytes, err := h.clubService.UploadProfileImage(clubID, reqImage)
	if err != nil {
		return err
	}

	return c.Blob(200, http.DetectContentType(fileBytes), fileBytes)
}

func (h *ClubHandler) GetProfileImage(c echo.Context) error {
	clubID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.Wrap(err, "알 수 없는 오류가 발생했습니다.")
	}

	fileBytes, err := h.clubService.GetProfileImage(clubID)
	if err != nil {
		return err
	}

	return c.Blob(200, http.DetectContentType(fileBytes), fileBytes)
}
