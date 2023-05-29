package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// set up kafka consumer congfig
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",         //broker address
		"group.id":          "console-consumer-21564", //consumer group
		"auto.offset.reset": "earliest",               //earliest offset
	}

	// cratte kafka consume

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(err)
	}

	// subscribe to target topic

	consumer.SubscribeTopics([]string{"mytopic"}, nil)

	// Consume messages

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Recived message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error while consuming message: %v (%v)\n", err, msg)
		}
	}

	// close kafka consumer
	consumer.Close()

}
