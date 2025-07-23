package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	ctx := context.Background()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "topic-1",
	})
	defer writer.Close()

	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte("key1"),
		Value: []byte("Some message"),
	})
	if err != nil {
		log.Fatalln("Не удалось отправить соо")
	}

	select {
	case <-ctx.Done():
		log.Println("Context  был прерван")
	}
}
