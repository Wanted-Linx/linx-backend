package service

import (
	"github.com/Wanted-Linx/linx-backend/api/domain"
	"github.com/Wanted-Linx/linx-backend/api/ent"
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

func (s *companyService) GetCompanyByID(studentID int) (*domain.CompanyDto, error) {

	return nil, nil
}
