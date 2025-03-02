package usecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"xm-test-ilya-chicherin/entity/company"
	"xm-test-ilya-chicherin/repository"
	"xm-test-ilya-chicherin/services/company/v1/models"
)

type CompanyUC struct {
	kafkaRepository   *repository.KafkaRepository
	companyRepository *repository.CompaniesRepository
}

func NewCompany(kafkaWriter *repository.KafkaRepository) *CompanyUC {
	companyRepo := repository.NewRepository()
	return &CompanyUC{
		companyRepository: companyRepo,
		kafkaRepository:   kafkaWriter,
	}
}

func (c *CompanyUC) GetCompany(ctx context.Context, companyID string) (company.Company, error) {
	dbCompany, err := c.companyRepository.ReadCompany(ctx, companyID)
	if err != nil {
		return company.Company{}, err
	}
	return dbCompany.ToEntity(), nil
}

func (c *CompanyUC) DeleteCompany(ctx context.Context, companyID string) error {
	err := c.kafkaRepository.PublishDeleteEvent(ctx, companyID)
	if err != nil {
		return err
	}
	return nil
}

func (c *CompanyUC) CreateCompany(ctx context.Context, req company.Company) error {
	req.ID = uuid.New()
	err := c.kafkaRepository.PublishCreateEvent(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func (c *CompanyUC) ProcessCreateEvent(ctx context.Context, newCompany company.Company) error {
	log.Printf("processing create event for company: %s", newCompany.ID)
	dbCompany, err := c.companyRepository.ReadCompanyByName(ctx, newCompany.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("failed to insert company %s: %v", newCompany.ID, err)
		return err
	}
	if dbCompany.Name == newCompany.Name {
		return nil
	}
	companyModel := models.FromEntity(newCompany)
	err = c.companyRepository.InsertCompany(ctx, companyModel)
	if err != nil {
		log.Printf("failed to insert company %s: %v", newCompany.ID, err)
		return err
	}
	return nil
}

func (c *CompanyUC) UpdateCompany(ctx context.Context, companyID string, updates map[string]interface{}) error {
	err := c.kafkaRepository.PublishUpdateEvent(ctx, companyID, updates)
	if err != nil {
		log.Printf("failed to publish update event: %v", err)
		return err
	}

	return nil
}

func (c *CompanyUC) ProcessUpdateEvent(ctx context.Context, companyID string, updates map[string]interface{}) error {
	err := c.companyRepository.UpdateCompany(ctx, companyID, updates)
	if err != nil {
		log.Printf("failed to update company %s: %v", companyID, err)
		return err
	}
	return nil
}
