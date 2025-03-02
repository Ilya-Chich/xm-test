package company

import (
	"github.com/google/uuid"
)

type Type string

const (
	Corporations       Type = "Corporations"
	NonProfit          Type = "NonProfit"
	Cooperative        Type = "Cooperative"
	SoleProprietorship Type = "Sole Proprietorship"
)

type Company struct {
	ID                uuid.UUID
	Name              string
	Description       *string
	AmountOfEmployees int
	Registered        bool
	Type              Type
}
