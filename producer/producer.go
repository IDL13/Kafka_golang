package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	leader, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
	if err != nil {
		log.Println("Eror in leader connection")
	}
	leader.SetWriteDeadline(time.Now().Add(time.Second * 10))
	leader.WriteMessages(kafka.Message{Value: []byte("from golang message")})
}
