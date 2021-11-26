package service

import (
	"time"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type projectClubService struct {
	projectClubRepo domain.ProjectClubRepository
}

func NewProjectClubService(projectClubRepo domain.ProjectClubRepository) domain.ProjectClubService {
	return &projectClubService{projectClubRepo: projectClubRepo}
}

func (s *projectClubService) ApplyProjectClub(reqPc *domain.ProjectClubApplyRequest) error {

	// default로 일단 무조건 현재 시간으로...
	reqPc.StartDate = util.ConvertTimeToStr(time.Now())
	log.Println(reqPc.StartDate)
	projectClub, err := s.projectClubRepo.Apply(reqPc.ClubID, reqPc.ProjectID, reqPc.StartDate)
	if err != nil {
		return errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}
	log.Info("프로젝트 신청 완료", projectClub)
	return nil
}
