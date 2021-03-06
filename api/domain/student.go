package domain

import (
	"log"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type StudentDto struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	University     string        `json:"university"`
	Description    *string       `json:"description"`
	InterestedType []string      `json:"interested_type"`
	ProfileLink    *string       `json:"profile_link"`
	ProfileImage   *string       `json:"profile_image"`
	Clubs          []*JoinedClub `json:"clubs"`
}

// type StudentProfileImageDto struct {
// 	Image []byte `json:"image"`
// }

// Multipart form-data로 이미지랑 같이 받을까...?
type StudentProfileUpdate struct {
	Description    string   `json:"description"`
	InterestedType []string `json:"interested_type"`
	ProfileLink    string   `json:"profile_link"`
}

// type StudentProfileImage struct {
// 	Image []*multipart.FileHeader `json:"image"`
// }

type StudentService interface {
	Save(userID int, reqSignup *SignUpRequest) (*StudentDto, error)
	GetStudentByID(studentID int) (*StudentDto, error)
	UpdateProfile(studentID int, reqStudent *StudentProfileUpdate) (*StudentDto, error)
	UploadProfileImage(studentID int, reqImage *ProfileImageRequest) ([]byte, error)
	GetProfileImage(studentID int) ([]byte, error)
}

type StudentRepository interface {
	Save(reqStudent *ent.Student) (*ent.Student, error)
	GetByID(studentID int, reqStudent *ent.Student) (*ent.Student, error)
	UpdateProfile(reqStudent *ent.Student) (*ent.Student, error)
	UploadProfileImage(reqStudent *ent.Student) (*ent.Student, error)
	GetAllTasks(studentID int) ([]*ent.TaskType, error)
	SaveTasks(s *ent.Student, taskType *ent.TaskType) (*ent.TaskType, error)
	// GetProfileImage(reqStudent *ent.Student) ([]byte, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func StudentToDto(src *ent.Student, srcJoinedClubs []*ent.Club) *StudentDto {
	// clubs, _ := src.QueryClubMember().QueryClub().All(context.Background())
	joinedclubsDto := MemberClubsToDto(src.Edges.Club)
	interestedTypes := make([]string, len(src.Edges.TaskType))
	for i := 0; i < len(src.Edges.TaskType); i++ {
		interestedTypes[i] = src.Edges.TaskType[i].Type
	}

	log.Println(interestedTypes, src.Edges.TaskType)
	return &StudentDto{
		ID:           src.ID,
		Name:         src.Name,
		University:   src.University,
		Description:  src.Description,
		ProfileLink:  src.ProfileLink,
		ProfileImage: src.ProfileImage,
		Clubs:        joinedclubsDto,
		// TODO: interested_type db에서 가져오기(지금은 임시로...)
		InterestedType: interestedTypes,
	}
}
