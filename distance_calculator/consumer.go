package main

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/lucaliebenberg/tolling/types"
	"github.com/sirupsen/logrus"
)

// This can also be called a KafkaTransporter
type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, svc CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}
	c.SubscribeTopics([]string{topic}, nil)
	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka transport started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {

	for c.isRunning {
		msg, err := c.consumer.ReadMessage(-1) // could set to every second -> (time.Second)
		if err != nil {
			logrus.Errorf("kafka consume error %s", err)
			continue
		}
		var data types.OBUData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error %s", err)
			continue
		}
		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Service calculation error %s", err)
			continue
		}
		fmt.Printf("distance %.2f\n", distance)
	}
}
