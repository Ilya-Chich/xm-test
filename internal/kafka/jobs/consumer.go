package repository

import (
	"xm-test-ilya-chicherin/repository"
	"xm-test-ilya-chicherin/usecase"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader            *kafka.Reader
	companyUC         *usecase.CompanyUC
	companyRepository *repository.CompaniesRepository
}

func NewKafkaConsumer(brokerAddress, topic string, uc *usecase.CompanyUC, companyRepo *repository.CompaniesRepository) *KafkaConsumer {
	return &KafkaConsumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{brokerAddress},
			Topic:    topic,
			GroupID:  "delete-company-group",
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		}),
		companyUC:         uc,
		companyRepository: companyRepo,
	}
}
