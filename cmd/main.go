package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dungly116/order-processing/internal/kafkaconfig"
	"github.com/dungly116/order-processing/internal/orderprocessor"
	"github.com/dungly116/order-processing/internal/orderreceiver"
)

func main() {
	kafkaConfig, err := kafkaconfig.LoadKafkaConfig("kafka-config.json")
	if err != nil {
		panic("Error loading Kafka config: " + err.Error())
	}

	orderProducer := orderreceiver.NewOrderProducer(kafkaConfig)

	go func() {
		for {
			orderreceiver.OrderReceivingModule(orderProducer)
			time.Sleep(time.Second)
		}
	}()

	go orderprocessor.KafkaConsumerModule(kafkaConfig)

	waitForInterrupt()
	orderProducer.Close()
}

func waitForInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
