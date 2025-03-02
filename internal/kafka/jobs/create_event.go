package repository

import (
	"context"
	"log"

	"xm-test-ilya-chicherin/entity/company"
	"xm-test-ilya-chicherin/pkg/json"
)

func (kc *KafkaConsumer) StartCreateConsumer(ctx context.Context) {
	go func() {
		for {
			msg, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error reading create event from Kafka: %v", err)
				continue
			}

			var event struct {
				EventType string          `json:"event_type"`
				Company   company.Company `json:"company"`
			}

			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("error unmarshaling Kafka event: %v", err)
				continue
			}
			if event.EventType == "create" {
				err := kc.companyUC.ProcessCreateEvent(ctx, event.Company)
				if err != nil {
					log.Printf("Failed to insert company %s: %v", event.Company.ID, err)
				}
			}
		}
	}()
}
