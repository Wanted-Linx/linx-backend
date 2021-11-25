package service

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

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
func (s *companyService) UploadProfileImage(companyID int, reqImage *domain.CompanyProfileImage) ([]byte, error) {
	var fileBytes []byte

	// TODO: profile upload 함수 공통이므로 따로 빼기...
	for _, imageFile := range reqImage.Image {
		// Source
		src, err := imageFile.Open()
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}
		defer src.Close()

		fileBytes, err = ioutil.ReadAll(src)
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}

		// Destination
		dir := fmt.Sprintf("./companies/profile/%d/image", companyID)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0700) // Create your file
		}

		// image 저장 위치 uuid 값을 key로...
		key := uuid.New().String()
		dst, err := os.Create(filepath.Join(dir, filepath.Base(key)))
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}
		defer dst.Close()

		img, fileType, err := image.Decode(bytes.NewReader(fileBytes))
		if err != nil {
			return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
		}

		switch fileType {
		case "jpeg":
			log.Println(fileType)
			err = jpeg.Encode(dst, img, nil)
			if err != nil {
				return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
			}
		default:
			err = png.Encode(dst, img)
			if err != nil {
				return nil, errors.Wrap(err, "프로필 이미지 업로드 실패")
			}
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
	}

	return fileBytes, nil
}

func (s *companyService) GetProfileImage(companyID int) ([]byte, error) {
	company, err := s.GetCompanyByID(companyID)
	if err != nil {
		return nil, err
	}

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("./companies/profile/%d/image/%s", companyID, *company.ProfileImage))
	if err != nil {
		return nil, errors.Wrap(err, "프로필 이미지 조회 실패")
	}

	log.Info("프로필 이미지 조회 완료", company)
	return fileBytes, nil
}
