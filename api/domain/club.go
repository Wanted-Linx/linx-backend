package domain

import (
	"mime/multipart"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type ClubDto struct {
	ID             int            `json:"id"`
	LeaderID       int            `json:"leader_id"`
	LeaderName     string         `json:"leader_name"`
	Name           string         `json:"name"`
	Organization   string         `json:"organization"`
	Description    string         `json:"description"`
	InterestedType []string       `json:"interested_type"`
	ProfileLink    *string        `json:"profile_link"`
	ProfileImage   *string        `json:"profile_image"` // image path 저장
	ClubMembers    []*ClubMember  `json:"club_members"`  // 클럽에 가입한 멤버들
	ClubProjects   []*ClubProject `json:"club_projects"` // 클럽이 진행하는 프로젝트의 로그들도 같이 있음
}

type ClubCreateRequest struct {
	Name           string                  `json:"name"`
	Organization   string                  `json:"organization"`
	Description    string                  `json:"description"`
	InterestedType []string                `json:"interested_type"`
	ProfileLink    string                  `json:"profile_link"`
	ProfileImage   []*multipart.FileHeader `json:"profile_image"`
}

type ClubService interface {
	CreateClub(studentID int, reqClub *ClubCreateRequest) (*ClubDto, error)
	GetClubByID(clubID int) (*ClubDto, error)
	GetAllClubs(limit, offset int) ([]*ClubDto, error)
}

type ClubRepository interface {
	Save(reqClub *ent.Club) (*ent.Club, error)
	GetByID(clubID int) (*ent.Club, []*ent.Student, error)
	GetAll(limit, offset int) ([]*ent.Club, [][]*ent.Student, error)
}

func ClubToDto(srcClub *ent.Club, srcClubMembers []*ent.Student) *ClubDto {
	clubMembersDto := ClubMembersToDto(srcClubMembers)
	clubProjectsDto := ClubProjectsToDto(srcClub.Edges.Project)

	return &ClubDto{
		ID:           srcClub.ID,
		LeaderID:     srcClub.Edges.Leader.ID,
		LeaderName:   srcClub.Edges.Leader.Name,
		Name:         srcClub.Name,
		Organization: srcClub.Organization,
		Description:  srcClub.Description,
		ProfileLink:  srcClub.ProfileLink,
		ProfileImage: srcClub.ProfileImage,
		ClubMembers:  clubMembersDto,
		ClubProjects: clubProjectsDto,
	}
}

//club.QueryClubMember().QueryStudent().All()
