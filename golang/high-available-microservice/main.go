

package main

import (
    "fmt"
    "log"
    "time"

    "github.com/confluentinc/confluent-kafka-go/kafka"
)


func main() {
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "acks": "all",
        "retries": 0,
    })
    if err != nil {
        log.Fatal(err)
    }
    defer producer.Close()

    topic := "my-topic"
    for i := 0; i < 10; i++ {
        value := fmt.Sprintf("Hello Go! %d", i)
        producer.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value: []byte(value),
        }, nil)
        time.Sleep(1 * time.Second)
    }
}


