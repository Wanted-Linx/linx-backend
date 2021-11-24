package service

import (
	"context"
	"fmt"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type clubService struct {
	clubRepo       domain.ClubRepository
	clubMemberRepo domain.ClubMemberRepository
}

func NewClubService(clubRepo domain.ClubRepository, clubMemberRepo domain.ClubMemberRepository) domain.ClubService {
	return &clubService{
		clubRepo:       clubRepo,
		clubMemberRepo: clubMemberRepo,
	}
}

func (s *clubService) CreateClub(clubLeaderID int, reqClub *domain.ClubCreateRequest) (*domain.ClubDto, error) {

	// TODO: profile image 먼저 upload 후 저장된 path 값을 ProfileImage에 집어넣는다.
	// 우선은 default로 '/profile/club/clubLeaderID/
	profile_key := fmt.Sprintf("/profile/club/%d", clubLeaderID)

	club := &ent.Club{
		Name:         reqClub.Name,
		Organization: reqClub.Organization,
		Description:  reqClub.Description,
		ProfileLink:  &reqClub.ProfileLink,
		ProfileImage: &profile_key,
		Edges: ent.ClubEdges{
			Leader: &ent.Student{ID: clubLeaderID},
		},
	}

	newClub, err := s.clubRepo.Save(club)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	log.Println(newClub, clubLeaderID, newClub.Edges.Leader)
	_, err = s.clubMemberRepo.Register(clubLeaderID, newClub.ID)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}
	clubMembers, err := newClub.QueryClubMember().QueryStudent().All(context.Background())
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("동아리 생성 완료", newClub, clubMembers)
	return domain.ClubToDto(newClub, clubMembers), nil
}

func (s *clubService) GetClubByID(clubID int) (*domain.ClubDto, error) {

	return nil, nil
}
