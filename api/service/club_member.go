package service

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type clubMemberService struct {
	clubMemberRepo domain.ClubMemberRepository
}

func NewClubMemberService(clubMemberRepo domain.ClubMemberRepository) domain.ClubMemberService {
	return &clubMemberService{clubMemberRepo: clubMemberRepo}
}

func (s *clubMemberService) RegisterClubMember(studentID, clubID int) error {
	clubMember, err := s.clubMemberRepo.Register(studentID, clubID)
	if err != nil {
		return errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}
	log.Info("동아리 가입 완료", clubMember)
	return nil
}
