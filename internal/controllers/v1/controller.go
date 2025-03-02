package v1

import "xm-test-ilya-chicherin/usecase"

type CompanyController struct {
	buildVersion string
	CompanyUC    usecase.Company
}

func NewCompanyController() (*CompanyController, error) {
	return &CompanyController{}, nil
}
