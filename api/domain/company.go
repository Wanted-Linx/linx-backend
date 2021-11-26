package domain

import (
	"mime/multipart"

	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type CompanyDto struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	BusinessType   []string `json:"business_type"`
	Description    *string  `json:"description"`
	BusinessNumber string   `json:"business_number"`
	ProfileImage   *string  `json:"profile_image"`
	Homepage       *string  `json:"hompage"`
}

type CompanyProfileImageDto struct {
	Image []byte `json:"image"`
}

// Multipart form-data로 이미지랑 같이 받을까...?
type CompanyProfileUpdate struct {
	BusinessType []string `json:"business_type"`
	Homepage     string   `json:"homepage"`
	Description  string   `json:"description"`
}

type CompanyProfileImage struct {
	Image []*multipart.FileHeader `json:"image"`
}

type CompanyService interface {
	Save(userID int, reqSignup *SignUpRequest) (*CompanyDto, error)
	GetCompanyByID(companyID int) (*CompanyDto, error)
	GetAllCompanies(limit, offset int) ([]*CompanyDto, error)
	UpdateProfile(companyID int, reqCompany *CompanyProfileUpdate) (*CompanyDto, error)
	UploadProfileImage(companyID int, reqImage *CompanyProfileImage) ([]byte, error)
	GetProfileImage(companyID int) ([]byte, error)
}

type CompanyRepository interface {
	Save(reqStudent *ent.Company) (*ent.Company, error)
	GetByID(studentID int, reqStudent *ent.Company) (*ent.Company, error)
	GetAll(limit, offset int) ([]*ent.Company, error)
	UpdateProfile(reqCompany *ent.Company) (*ent.Company, error)
	UploadProfileImage(reqCompany *ent.Company) (*ent.Company, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func CompanyToDto(src *ent.Company) *CompanyDto {
	return &CompanyDto{
		ID:             src.ID,
		Name:           src.Name,
		BusinessNumber: src.BusinessNumber,
		Description:    src.Description,
		ProfileImage:   src.ProfileImage,
		Homepage:       src.Hompage,
		BusinessType:   []string{"개발", "디자인"}, // TODO: default로 이렇게 잡아뒀지만 추후 businesstype 테이블 생성 후 그에 맞게 수정...
	}
}

func CompaniesToDto(src []*ent.Company) []*CompanyDto {
	companiesDto := make([]*CompanyDto, 0)

	for _, company := range src {
		companyDto := CompanyToDto(company)
		companiesDto = append(companiesDto, companyDto)
	}

	return companiesDto
}
