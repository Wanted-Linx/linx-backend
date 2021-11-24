package domain

import "github.com/Wanted-Linx/linx-backend/api/ent"

type CompanyDto struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BusinessType   []string `json:"business_type"`
	Description    *string  `json:"description"`
	BusinessNumber *string  `json:"business_number"`
	ProfileImage   *string  `json:"profile_image"`
}

type CompanyService interface {
	Save(userID int, reqSignup *SignUpRequest) (*CompanyDto, error)
	GetCompanyByID(companyID int) (*CompanyDto, error)
}
type CompanyRepository interface {
	Save(reqStudent *ent.Company) (*ent.Company, error)
	GetByID(studentID int, reqStudent *ent.Company) (*ent.Company, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func CompanyToDto(src *ent.Company) *CompanyDto {
	return &CompanyDto{
		ID:             src.Edges.User.ID,
		Name:           src.Name,
		BusinessNumber: &src.BusinessNumber,
		Description:    src.Description,
		ProfileImage:   src.ProfileImage,
		// BusinessType:  ,
	}
}
