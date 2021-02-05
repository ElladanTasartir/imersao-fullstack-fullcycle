package main

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChannel := make(chan kafka.Event)
	producer := NewKafkaProducer()

	Publish("Hello from Imersao FullStack && FullCycle", "imersao-desafio-2", producer, deliveryChannel)

	DeliveryReport(deliveryChannel)
}
