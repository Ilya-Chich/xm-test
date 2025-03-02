package repository

import (
	"context"
	"encoding/json"
	"log"
)

func (kc *KafkaConsumer) StartUpdateConsumer(ctx context.Context) {
	go func() {
		for {
			m, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error reading update event from Kafka: %v", err)
				continue
			}

			var event struct {
				EventType string                 `json:"event_type"`
				CompanyID string                 `json:"company_id"`
				Updates   map[string]interface{} `json:"updates"`
			}

			if err := json.Unmarshal(m.Value, &event); err != nil {
				log.Printf("error unmarshaling Kafka update event: %v", err)
				continue
			}

			if event.EventType == "update" {
				err := kc.companyUC.ProcessUpdateEvent(ctx, event.CompanyID, event.Updates)
				if err != nil {
					log.Printf("Failed to update company %s: %v", event.CompanyID, err)
				}
			}
		}
	}()
}
