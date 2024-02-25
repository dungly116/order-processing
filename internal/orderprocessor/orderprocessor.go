package orderprocessor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dungly116/order-processing/internal/kafkaconfig"
	"github.com/dungly116/order-processing/models"
	"github.com/segmentio/kafka-go"
)

func KafkaConsumerModule(config kafkaconfig.Config) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config.Brokers,
		Topic:     config.Topic,
		Partition: 0,
		MinBytes:  10e3,
		MaxBytes:  10e6,
	})

	defer r.Close()

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		order := models.Order{}
		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Println("Error decoding order:", err)
			continue
		}

		fmt.Println("Processing Order:", order)
	}
}
