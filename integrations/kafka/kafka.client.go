// https://docs.confluent.io/kafka-clients/go/current/overview.html
package kafka_client

import (
	"fmt"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	producer *kafka.Producer
}

func createProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "myProducer",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	return p
}

// start: singleton logic

var instance *KafkaClient
var once sync.Once

func NewKafkaClient() *KafkaClient {
	once.Do(func() {
		instance = &KafkaClient{
			producer: createProducer(),
		}
	})
	return instance
}

// end: singleton logic

func (k *KafkaClient) Produce(message string, topic string) {
	deliveryChan := make(chan kafka.Event)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Failed to produce message: %v\n", err)
		close(deliveryChan)
		return
	}

	event := <-deliveryChan
	messageCreated := event.(*kafka.Message)

	if messageCreated.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", messageCreated.TopicPartition.Error)
	}

	close(deliveryChan)
}
