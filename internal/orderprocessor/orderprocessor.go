package orderprocessor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dungly116/order-processing/internal/kafkaconfig"
	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID       string  `json:"id"`
	Product  string  `json:"product"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

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

		order := Order{}
		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Println("Error decoding order:", err)
			continue
		}

		// Process the order and store in the database sequentially.
		fmt.Println("Processing Order:", order)
		// Add your processing logic and database storage here.
	}
}
