package service

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type clubService struct {
	clubRepo       domain.ClubRepository
	clubMemberRepo domain.ClubMemberRepository
	taskTypeRepo   domain.TaskTypeRepository
}

func NewClubService(clubRepo domain.ClubRepository,
	clubMemberRepo domain.ClubMemberRepository,
	taskTypeRepo domain.TaskTypeRepository) domain.ClubService {
	return &clubService{
		clubRepo:       clubRepo,
		clubMemberRepo: clubMemberRepo,
		taskTypeRepo:   taskTypeRepo,
	}
}

func (s *clubService) CreateClub(clubLeaderID int, reqClub *domain.ClubCreateRequest) (*domain.ClubDto, error) {
	club := &ent.Club{
		Name:         reqClub.Name,
		Organization: reqClub.Organization,
		Description:  reqClub.Description,
		ProfileLink:  &reqClub.ProfileLink,
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

	tasks := make([]*ent.TaskType, len(reqClub.InterestedType))
	for i := 0; i < len(reqClub.InterestedType); i++ {
		interestedType := &ent.TaskType{
			Type: reqClub.InterestedType[i],
			Edges: ent.TaskTypeEdges{
				Club: &ent.Club{ID: newClub.ID},
			},
		}
		task, err := s.clubRepo.SaveTasks(newClub, interestedType)
		if err != nil {
			return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
		}

		tasks[i] = task
	}

	newClub.Edges.TaskType = tasks
	log.Info("동아리 생성 완료", newClub, newClub.Edges.TaskType)
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
	tasks, err := s.clubRepo.GetAllTasks(club.ID)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	club.Edges.TaskType = tasks

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
		tasks, err := s.clubRepo.GetAllTasks(club.ID)
		if err != nil {
			return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
		}
		club.Edges.TaskType = tasks
		clubDto := domain.ClubToDto(club, allMembers[idx])
		clubsDto[idx] = clubDto
	}

	return clubsDto, nil
}

func (s *clubService) UploadProfileImage(clubID int, reqImage *domain.ProfileImageRequest) ([]byte, error) {
	// Destination
	dir := fmt.Sprintf("./clubs/profile/%d/image", clubID)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700) // Create your file
		if err != nil {
			return nil, errors.Wrap(err, "이미지 저장용 디렉토리 생성 실패")
		}
	}

	key := uuid.New().String()
	fileBytes, err := UploadImage(clubID, dir, key, reqImage)
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
	}

	club := &ent.Club{
		ID:           clubID,
		ProfileImage: &key,
	}

	cup, err := s.clubRepo.UploadProfileImage(club)
	if err != nil {
		return nil, errors.WithMessage(err, "프로필 이미지 업로드 실패")
	}

	log.Info("프로필 이미지 업로드 성공", cup)
	return fileBytes, nil
}

func (s *clubService) GetProfileImage(clubID int) ([]byte, error) {
	club, err := s.GetClubByID(clubID)
	if err != nil {
		return nil, err
	}

	if club.ProfileImage == nil {
		return nil, errors.New("프로필 이미지가 존재하지 않습니다.")
	}

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./clubs/profile/%d/image/%s", clubID, *club.ProfileImage))
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 조회 실패")
	}

	return fileBytes, nil
}
