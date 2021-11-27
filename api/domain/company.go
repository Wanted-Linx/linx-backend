package domain

import (
	"github.com/Wanted-Linx/linx-backend/api/ent"
)

type CompanyDto struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Address        *string  `json:"address"`
	BusinessType   []string `json:"business_type"`
	Description    *string  `json:"description"`
	BusinessNumber string   `json:"business_number"`
	ProfileImage   *string  `json:"profile_image"`
	Homepage       *string  `json:"hompage"`
}

// type CompanyProfileImageDto struct {
// 	Image []byte `json:"image"`
// }

// Multipart form-data로 이미지랑 같이 받을까...?
type CompanyProfileUpdate struct {
	BusinessType []string `json:"business_type"`
	Homepage     string   `json:"homepage"`
	Description  string   `json:"description"`
	Address      string   `json:"address"`
}

// type CompanyProfileImage struct {
// 	Image []*multipart.FileHeader `json:"image"`
// }

type CompanyService interface {
	Save(userID int, reqSignup *SignUpRequest) (*CompanyDto, error)
	GetCompanyByID(companyID int) (*CompanyDto, error)
	GetAllCompanies(limit, offset int) ([]*CompanyDto, error)
	UpdateProfile(companyID int, reqCompany *CompanyProfileUpdate) (*CompanyDto, error)
	UploadProfileImage(companyID int, reqImage *ProfileImageRequest) ([]byte, error)
	GetProfileImage(companyID int) ([]byte, error)
}

type CompanyRepository interface {
	Save(reqStudent *ent.Company) (*ent.Company, error)
	GetByID(studentID int, reqStudent *ent.Company) (*ent.Company, error)
	GetAll(limit, offset int) ([]*ent.Company, error)
	UpdateProfile(reqCompany *ent.Company) (*ent.Company, error)
	UploadProfileImage(reqCompany *ent.Company) (*ent.Company, error)
	GetAllTasks(companyID int) ([]*ent.TaskType, error)
	SaveTasks(c *ent.Company, taskType *ent.TaskType) (*ent.TaskType, error)
	// GetAll(clubID int) ([]*ent.Student, error)
}

func CompanyToDto(src *ent.Company) *CompanyDto {
	taskTypes := make([]string, len(src.Edges.TaskType))
	for i := 0; i < len(src.Edges.TaskType); i++ {
		taskTypes[i] = src.Edges.TaskType[i].Type
	}

	return &CompanyDto{
		ID:             src.ID,
		Name:           src.Name,
		Address:        src.Address,
		BusinessNumber: src.BusinessNumber,
		Description:    src.Description,
		ProfileImage:   src.ProfileImage,
		Homepage:       src.Hompage,
		BusinessType:   taskTypes, // TODO: default로 이렇게 잡아뒀지만 추후 businesstype 테이블 생성 후 그에 맞게 수정...
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
