package repository

import (
	"context"
	"encoding/json"
	"log"
	cmp "xm-test-ilya-chicherin/entity/company"

	"github.com/segmentio/kafka-go"
)

const (
	DeletePartition = 0
	CreatePartition = 1
	PatchPartition  = 2
)

type KafkaRepository struct {
	writer *kafka.Writer
}

func NewKafkaRepository(writer *kafka.Writer) *KafkaRepository {
	return &KafkaRepository{writer: writer}
}

func (r *KafkaRepository) PublishDeleteEvent(ctx context.Context, companyID string) error {
	event := struct {
		EventType string `json:"event_type"`
		CompanyID string `json:"company_id"`
	}{
		EventType: "delete",
		CompanyID: companyID,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("error marshaling delete event for company %s: %v", companyID, err)
		return err
	}

	err = r.writer.WriteMessages(ctx, kafka.Message{
		Key:       []byte(companyID),
		Value:     message,
		Partition: DeletePartition,
	})
	if err != nil {
		log.Printf("error publishing delete event for company %s: %v", companyID, err)
		return err
	}

	log.Printf("successfully published delete event for company %s", companyID)
	return nil
}

func (r *KafkaRepository) PublishCreateEvent(ctx context.Context, company cmp.Company) error {
	event := struct {
		EventType string      `json:"event_type"`
		Company   cmp.Company `json:"company"`
	}{
		EventType: "create",
		Company:   company,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("error marshaling create event: %v", err)
		return err
	}

	err = r.writer.WriteMessages(ctx, kafka.Message{
		Key:       []byte(event.Company.ID.String()),
		Value:     message,
		Partition: CreatePartition,
	})
	if err != nil {
		log.Printf("error publishing create event: %v", err)
		return err
	}

	return nil
}

func (r *KafkaRepository) PublishUpdateEvent(ctx context.Context, companyID string, updates map[string]interface{}) error {
	event := struct {
		EventType string                 `json:"event_type"`
		CompanyID string                 `json:"company_id"`
		Updates   map[string]interface{} `json:"updates"`
	}{
		EventType: "update",
		CompanyID: companyID,
		Updates:   updates,
	}

	message, err := json.Marshal(event)
	if err != nil {
		log.Printf("error marshaling update event: %v", err)
		return err
	}

	err = r.writer.WriteMessages(ctx, kafka.Message{
		Key:       []byte(companyID),
		Value:     message,
		Partition: PatchPartition,
	})
	if err != nil {
		log.Printf("error publishing update event: %v", err)
		return err
	}
	return nil
}
