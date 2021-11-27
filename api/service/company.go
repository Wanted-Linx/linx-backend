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

type companyService struct {
	companyRepo domain.CompanyRepository
}

func NewCompanyService(companyRepo domain.CompanyRepository) domain.CompanyService {
	return &companyService{companyRepo: companyRepo}
}

func (s *companyService) Save(userID int, reqSignup *domain.SignUpRequest) (*domain.CompanyDto, error) {
	company := &ent.Company{
		ID:             userID,
		Name:           reqSignup.Name,
		BusinessNumber: reqSignup.BusinessNumber,
		Edges: ent.CompanyEdges{
			User: &ent.User{ID: userID},
		},
	}

	newCompany, err := s.companyRepo.Save(company)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	log.Info("회원가입(기업) 완료", newCompany)
	return domain.CompanyToDto(newCompany), nil
}

func (s *companyService) GetCompanyByID(companyID int) (*domain.CompanyDto, error) {
	company := &ent.Company{
		ID: companyID,
	}

	c, err := s.companyRepo.GetByID(companyID, company)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.WithMessage(err, "해당하는 기업을 찾을 수 없습니다.")
		}
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("해당하는 기업 조회 완료", c)
	return domain.CompanyToDto(c), nil
}

func (s *companyService) GetAllCompanies(limit, offset int) ([]*domain.CompanyDto, error) {
	companies, err := s.companyRepo.GetAll(limit, offset)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 오류가 발생했습니다.")
	}

	log.Info("기업들 조회 완료", companies)
	return domain.CompaniesToDto(companies), nil
}

func (s *companyService) UpdateProfile(companyID int, reqCompany *domain.CompanyProfileUpdate) (*domain.CompanyDto, error) {
	company := &ent.Company{
		ID:          companyID,
		Hompage:     &reqCompany.Homepage,
		Description: &reqCompany.Description,
	}

	// TODO: business type에 대한 처리도 추후에...

	updatedCompany, err := s.companyRepo.UpdateProfile(company)
	if err != nil {
		return nil, errors.WithMessage(err, "알 수 없는 에러가 발생했습니다.")
	}

	return domain.CompanyToDto(updatedCompany), nil
}
func (s *companyService) UploadProfileImage(companyID int, reqImage *domain.ProfileImageRequest) ([]byte, error) {
	dir := fmt.Sprintf("./companies/profile/%d/image", companyID)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700) // Create your file
		if err != nil {
			return nil, errors.Wrap(err, "이미지 저장용 디렉토리 생성 실패")
		}
	}

	key := uuid.New().String()
	fileBytes, err := UploadImage(companyID, dir, key, reqImage)
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
	}

	company := &ent.Company{
		ID:           companyID,
		ProfileImage: &key,
	}

	c, err := s.companyRepo.UploadProfileImage(company)
	if err != nil {
		return nil, errors.WithMessage(err, "프로필 이미지 업로드 실패")
	}

	log.Info("프로필 이미지 업로드 성공", c)

	return fileBytes, nil
}

func (s *companyService) GetProfileImage(companyID int) ([]byte, error) {
	company, err := s.GetCompanyByID(companyID)
	if err != nil {
		return nil, err
	}

	if company.ProfileImage == nil {
		return nil, errors.New("프로필 이미지가 존재하지 않습니다.")
	}

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./companies/profile/%d/image/%s", companyID, *company.ProfileImage))
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 조회 실패")
	}

	log.Info("프로필 이미지 조회 완료", company)
	return fileBytes, nil
}
