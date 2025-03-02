package repository

import (
	"context"
	"xm-test-ilya-chicherin/services/company/v1/models"
)

func (cr *CompaniesRepository) ReadCompany(ctx context.Context, id string) (models.Company, error) {
	var dbCompany models.Company
	err := cr.db.WithContext(ctx).Where("id = ?", id).First(&dbCompany).Error
	if err != nil {
		return models.Company{}, err
	}
	return dbCompany, nil
}

func (cr *CompaniesRepository) ReadCompanyByName(ctx context.Context, name string) (models.Company, error) {
	var dbCompany models.Company
	err := cr.db.WithContext(ctx).Where("name = ?", name).First(&dbCompany).Error
	if err != nil {
		return models.Company{}, err
	}
	return dbCompany, nil
}

func (cr *CompaniesRepository) DeleteCompany(companyID string) error {
	result := cr.db.Where("id = ?", companyID).Delete(&models.Company{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (cr *CompaniesRepository) InsertCompany(ctx context.Context, company models.Company) error {
	result := cr.db.WithContext(ctx).Create(&company)
	return result.Error
}

func (cr *CompaniesRepository) UpdateCompany(ctx context.Context, companyID string, updates map[string]interface{}) error {
	result := cr.db.WithContext(ctx).Model(&models.Company{}).Where("id = ?", companyID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
