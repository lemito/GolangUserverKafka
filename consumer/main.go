package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"}, Topic: "topic-1", GroupID: "groupid-1",
	})
	defer reader.Close()

	msg, err := reader.ReadMessage(context.Background())
	if err != nil {
		log.Fatalln("Ошибка соо: ", err)
	}

	fmt.Println(string(msg.Value))
}
