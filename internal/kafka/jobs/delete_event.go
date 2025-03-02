package repository

import (
	"golang.org/x/net/context"
	"log"
	"xm-test-ilya-chicherin/pkg/json"
)

func (kc *KafkaConsumer) StartDeleteConsumer(ctx context.Context) {
	go func() {
		for {
			m, err := kc.reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("error reading message from Kafka: %v", err)
				continue
			}

			var event struct {
				EventType string `json:"event_type"`
				CompanyID string `json:"company_id"`
			}

			if err := json.Unmarshal(m.Value, &event); err != nil {
				log.Printf("error unmarshaling Kafka event: %v", err)
				continue
			}

			if event.EventType == "delete" {
				err := kc.companyRepository.DeleteCompany(event.CompanyID)
				if err != nil {
					log.Printf("failed to delete company %s: %v", event.CompanyID, err)
				}
			}
		}
	}()
}
