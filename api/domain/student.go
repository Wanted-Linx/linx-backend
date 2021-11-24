package domain

import "github.com/Wanted-Linx/linx-backend/api/ent"

type StudentDto struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	University     string  `json:"university"`
	InterestedType *string `json:"interested_type"`
	ProfileLink    *string `json:"profile_link"`
	ProfileImage   *string `json:"profile_image"`
}

type Student struct {
}
type StudentService interface {
	Save(userID int, reqSignup *SignUpRequest) (*StudentDto, error)
	GetStudentByID(studentID int) (*StudentDto, error)
}
type StudentRepository interface {
	Save(reqStudent *ent.Student) (*ent.Student, error)
	GetByID(studentID int, reqStudent *ent.Student) (*ent.Student, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func StudentToDto(src *ent.Student) *StudentDto {
	return &StudentDto{
		ID:           src.Edges.User.ID,
		Name:         src.Name,
		University:   src.University,
		ProfileLink:  src.ProfileLink,
		ProfileImage: src.ProfileImage,
	}
}
