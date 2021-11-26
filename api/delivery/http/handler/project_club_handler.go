package handler

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ProjectClubHandler struct {
	projectClubService domain.ProjectClubService
}

func NewProjectClubHandler(projectClubService domain.ProjectClubService) *ProjectClubHandler {
	return &ProjectClubHandler{projectClubService: projectClubService}
}

func (h *ProjectClubHandler) ApplyProjectClub(c echo.Context) error {
	// 인증용
	studentID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청 유저 id: ", studentID)
	reqPc := new(domain.ProjectClubApplyRequest)
	if err := c.Bind(reqPc); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	err = h.projectClubService.ApplyProjectClub(reqPc)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"result": "프로젝트 신청 성공"})
}
