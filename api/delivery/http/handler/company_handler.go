package handler

import "github.com/Wanted-Linx/linx-backend/api/domain"

type CompanyHandler struct {
	companySerivce domain.CompanyService
}

func NewCompanyHandler(companySerivce domain.CompanyService) *CompanyHandler {
	return &CompanyHandler{companySerivce: companySerivce}
}
