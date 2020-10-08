package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/dapr/go-sdk/service/http"
)

func main() {
	runSub := flag.Bool("sub", false, "Run subscriber instead of publisher")
	flag.Parse()

	ctx := context.Background()
	topicName := "foo"

	// pubsubName will be the value metadata.name from
	// deploy/stan-component.yaml.
	pubsubName := os.Getenv("DAPR_PUBSUB_NAME")

	var err error
	if *runSub {
		err = runSubscriber(ctx, pubsubName, topicName)
	} else {
		err = runPublisher(ctx, pubsubName, topicName)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func runPublisher(ctx context.Context, pubsubName, topicName string) error {
	log.Println("Starting publisher...")

	c, err := client.NewClient()
	if err != nil {
		return err
	}
	defer c.Close()

	var i int
	for range time.NewTicker(1 * time.Second).C {
		i++
		data := []byte(fmt.Sprintf("Hello %d!", i))

		err := c.PublishEvent(ctx, pubsubName, topicName, data)
		if err != nil {
			log.Println("publish on", topicName, "failed:", err)
			continue
		}
		log.Println("publish on", topicName, "ok")
	}

	return nil
}

func runSubscriber(ctx context.Context, pubsubName, topicName string) error {
	log.Println("Starting subscriber...")

	sub := &common.Subscription{
		PubsubName: pubsubName,
		Topic:      topicName,
		Route:      fmt.Sprintf("/%s", topicName),
	}
	log.Printf("subscription: %+v\n", sub)

	s := http.NewService(":3000")
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		return err
	}

	return s.Start()
}

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("eventHandler: %+v\n", e)
	return false, nil
}
