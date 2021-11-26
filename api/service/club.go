package service

import (
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
	// 동아리 생성한 사람은 기본적으로 동아리에 가입이 되어야 함(leader)
	_, err = s.clubMemberRepo.Register(clubLeaderID, newClub.ID)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	// club.QueryProjectClub().QueryProjects().All
	log.Info("동아리 생성 완료", newClub)
	return domain.ClubToDto(newClub, nil), nil
}

func (s *clubService) GetClubByID(clubID int) (*domain.ClubDto, error) {
	club, members, err := s.clubRepo.GetByID(clubID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "해당하는 동아리를 찾지 못했습니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("동아리 조회 완료", club, members)
	return domain.ClubToDto(club, members), nil
}

func (s *clubService) GetAllClubs(limit, offset int) ([]*domain.ClubDto, error) {
	clubs, allMembers, err := s.clubRepo.GetAll(limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	clubsDto := make([]*domain.ClubDto, len(clubs))

	for idx, club := range clubs {
		clubDto := domain.ClubToDto(club, allMembers[idx])
		clubsDto[idx] = clubDto
	}

	return clubsDto, nil
}
