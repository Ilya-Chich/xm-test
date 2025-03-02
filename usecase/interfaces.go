package usecase

import (
	"context"

	"xm-test-ilya-chicherin/entity/company"
)

type Company interface {
	GetCompany(ctx context.Context, companyID string) (company.Company, error)
	DeleteCompany(ctx context.Context, companyID string) error
	CreateCompany(ctx context.Context, company company.Company) error
	UpdateCompany(ctx context.Context, companyID string, companyUpdates map[string]interface{}) error
}

type Auth interface {
	Login() error
}
