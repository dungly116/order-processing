package orderreceiver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dungly116/order-processing/internal/kafkaconfig"
	"github.com/dungly116/order-processing/models"
	"github.com/segmentio/kafka-go"
)

func OrderReceivingModule(producer *kafka.Writer) {
	order := models.Order{ID: "123", Product: "ExampleProduct", Quantity: 2, Price: 19.99}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	message := kafka.Message{
		Value: orderJSON,
	}

	err = producer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Order sent:", order)
}

func NewOrderProducer(config kafkaconfig.Config) *kafka.Writer {
	writerConfig := kafka.WriterConfig{
		Brokers:  config.Brokers,
		Topic:    config.Topic,
		Balancer: &kafka.LeastBytes{},
	}

	// Initialize Kafka writer
	writer := kafka.NewWriter(writerConfig)

	return writer
}
