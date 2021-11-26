package handler

import (
	"strconv"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ProjectHandler struct {
	projectService domain.ProjectService
}

func NewProjectHandler(projectService domain.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

func (h *ProjectHandler) CreateProject(c echo.Context) error {
	companyID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	reqProject := new(domain.ProjectCreateRequest)
	if err := c.Bind(reqProject); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	newProject, err := h.projectService.CreateProject(companyID, reqProject)
	if err != nil {
		return err
	}

	return c.JSON(200, newProject)
}

func (h *ProjectHandler) GetProjectByID(c echo.Context) error {
	// 인증용
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청한 유저 id:", userID)
	projectID, err := strconv.Atoi(util.GetParams("project_id", c))
	if err != nil {
		return errors.WithStack(err)
	}

	project, err := h.projectService.GetProjectByID(projectID)
	if err != nil {
		return err
	}

	return c.JSON(200, project)
}

func (h *ProjectHandler) GetAllProjects(c echo.Context) error {
	// 인증용
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

	projects, err := h.projectService.GetAllProjects(limit, offset)
	if err != nil {
		return err
	}

	return c.JSON(200, projects)
}

func (h *ProjectHandler) CreateProjectLog(c echo.Context) error {
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청한 유저 id:", userID)

	reqPl := new(domain.ProjectLogCreateRequest)
	if err := c.Bind(reqPl); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	pl, err := h.projectService.CreateProjectLog(reqPl)
	if err != nil {
		return err
	}

	return c.JSON(200, pl)
}

func (h *ProjectHandler) CreateProjectLogFeedback(c echo.Context) error {
	userID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("조회 요청한 유저 id:", userID)

	reqPlf := new(domain.ProjectLogFeedbackRequest)
	if err := c.Bind(reqPlf); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	pl, err := h.projectService.CreateProjectLogFeedback(userID, reqPlf)
	if err != nil {
		return err
	}

	return c.JSON(200, pl)
}
