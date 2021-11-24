package handler

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
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
