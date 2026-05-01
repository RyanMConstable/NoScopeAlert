package main

import (
	"context"
	"encoding/json"
	"log"

	rmq "github.com/rabbitmq/rabbitmq-amqp-go-client/pkg/rabbitmqamqp"
)

const brokerURI = "amqp://guest:guest@noscopealert-rabbitmq:5672/"

type DownloadGameQueueItem struct {
	GameCode string `json:"game_code"`
	SteamID  string `json:"steam_id"`
}

func SendToQueue(gameCode string, steamid string) {
	ctx := context.Background()
	env := rmq.NewEnvironment(brokerURI, nil)
	conn, err := env.NewConnection(ctx)
	if err != nil {
		log.Panicf("Failed to connect to RabbitMQ: %v", err)
	}
	defer func() {
		_ = env.CloseConnections(context.Background())
	}()

	_, err = conn.Management().DeclareQueue(ctx, &rmq.QuorumQueueSpecification{Name: "download_game_queue"})
	if err != nil {
		log.Panicf("Failed to declare a queue: %v", err)
	}

	publisher, err := conn.NewPublisher(ctx, &rmq.QueueAddress{Queue: "download_game_queue"}, nil)
	if err != nil {
		log.Panicf("Failed to create publisher: %v", err)
	}
	defer func() { _ = publisher.Close(context.Background()) }()

	item := DownloadGameQueueItem{GameCode: gameCode, SteamID: steamid}
	body, err := json.Marshal(item)
	if err != nil {
		log.Panicf("Failed to marshal message: %v", err)
	}
	res, err := publisher.Publish(ctx, rmq.NewMessage([]byte(body)))
	if err != nil {
		log.Panicf("Failed to publish a message: %v", err)
	}
	switch res.Outcome.(type) {
	case *rmq.StateAccepted:
	default:
		log.Panicf("Unexpected publish outcome: %v", res.Outcome)
	}
	log.Printf(" [x] Sent %s", body)

}
