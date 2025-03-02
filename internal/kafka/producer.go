package kafka

import (
	"os"

	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:        os.Getenv("KAFKA_TOPIC"),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}
}
