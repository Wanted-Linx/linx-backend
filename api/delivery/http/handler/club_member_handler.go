package handler

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type ClubMemberHandler struct {
	clubMemberService domain.ClubMemberService
}

func NewClubMemberHandler(clubMembersService domain.ClubMemberService) *ClubMemberHandler {
	return &ClubMemberHandler{clubMemberService: clubMembersService}
}

func (h *ClubMemberHandler) RegisterClubMember(c echo.Context) error {
	studentID, err := util.GetRequestUserID(c)
	if err != nil {
		return errors.WithStack(err)
	}

	reqClubMemberReg := new(domain.ClubMemberRegisterRequest)
	if err := c.Bind(reqClubMemberReg); err != nil {
		return errors.Wrap(err, "잘못된 json body 입니다.")
	}

	err = h.clubMemberService.RegisterClubMember(studentID, reqClubMemberReg.ClubID)
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"result": "동아리 가입 성공"})
}
