package domain

import (
	"mime/multipart"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type ClubDto struct {
	ID             int           `json:"id"`
	LeaderID       int           `json:"leader_id"`
	Name           string        `json:"name"`
	Organization   string        `json:"organization"`
	Description    string        `json:"description"`
	InterestedType []string      `json:"interested_type"`
	ProfileLink    *string       `json:"profile_link"`
	ProfileImage   *string       `json:"profile_image"` // image path 저장
	ClubMembers    []*ClubMember `json:"club_members"`
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
}

type ClubRepository interface {
	Save(reqClub *ent.Club) (*ent.Club, error)
	GetByID(clubID int) (*ent.Club, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func ClubToDto(srcClub *ent.Club, srcClubMembers []*ent.Student) *ClubDto {
	clubMembersDto := ClubMembersToDto(srcClubMembers)

	return &ClubDto{
		ID:           srcClub.ID,
		LeaderID:     srcClub.Edges.Leader.ID,
		Name:         srcClub.Name,
		Organization: srcClub.Organization,
		Description:  srcClub.Description,
		ProfileLink:  srcClub.ProfileLink,
		ProfileImage: srcClub.ProfileImage,
		ClubMembers:  clubMembersDto,
	}
}

//club.QueryClubMember().QueryStudent().All()
