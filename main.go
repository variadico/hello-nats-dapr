package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dapr/go-sdk/client"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	ctx := context.Background()
	topicName := "foo"

	// pubsubName will be the value metadata.name from
	// deploy/stan-component.yaml.
	pubsubName := os.Getenv("DAPR_PUBSUB_NAME")

	var i int
	for range time.NewTicker(1 * time.Second).C {
		i++
		data := []byte(fmt.Sprintf("Hello %d!", i))

		err := c.PublishEvent(ctx, pubsubName, topicName, data)
		if err != nil {
			log.Println("publish failed:", err)
			continue
		}
		log.Println("publish ok")
	}
}
