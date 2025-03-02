package models

import (
	"github.com/google/uuid"
	"xm-test-ilya-chicherin/entity/company"
)

type CompanyType string

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID                uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name              string      `gorm:"type:varchar(15);unique;not null" json:"name"`
	Description       *string     `gorm:"type:varchar(3000)" json:"description,omitempty"`
	AmountOfEmployees int         `gorm:"not null" json:"amount_of_employees"`
	Registered        bool        `gorm:"not null" json:"registered"`
	Type              CompanyType `gorm:"type:varchar(20);not null;check:type IN ('Corporations', 'NonProfit', 'Cooperative', 'Sole Proprietorship')" json:"type"`
}

func (c *Company) ToEntity() company.Company {
	return company.Company{
		ID:                c.ID,
		Name:              c.Name,
		Description:       c.Description,
		AmountOfEmployees: c.AmountOfEmployees,
		Registered:        c.Registered,
		Type:              company.Type(c.Type),
	}
}

func FromEntity(c company.Company) Company {
	return Company{
		ID:                c.ID,
		Name:              c.Name,
		Description:       c.Description,
		AmountOfEmployees: c.AmountOfEmployees,
		Registered:        c.Registered,
		Type:              CompanyType(c.Type),
	}
}
