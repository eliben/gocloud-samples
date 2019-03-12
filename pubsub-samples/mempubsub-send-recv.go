package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gocloud.dev/pubsub"
	"gocloud.dev/pubsub/mempubsub"
)

func main() {
	// Open a topic and corresponding subscription.
	ctx := context.Background()
	t := mempubsub.NewTopic()
	defer t.Shutdown(ctx)
	s := mempubsub.NewSubscription(t, time.Second)
	defer s.Shutdown(ctx)

	// Send a message to the topic.
	if err := t.Send(ctx, &pubsub.Message{Body: []byte("Hello, world!")}); err != nil {
		log.Fatal(err)
	}

	// Receive a message from the subscription.
	m, err := s.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Print out the received message.
	fmt.Printf("%s\n", m.Body)

	// Acknowledge the message.
	m.Ack()
}
