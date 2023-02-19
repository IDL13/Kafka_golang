package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
	if err != nil {
		log.Println("Eror in conn connection")
	}
	conn.SetReadDeadline(time.Now().Add(time.Second * 8))

	messagesInKafka := conn.ReadBatch(1e3, 1e9)
	bytes := make([]byte, 1e3)

	for {
		_, err := messagesInKafka.Read(bytes)
		if err != nil {
			break
		}
		fmt.Println(string(bytes))
	}
}
