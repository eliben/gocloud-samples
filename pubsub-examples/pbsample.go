package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"
)

// To publish a message here, post a message to the subscription via the GCP
// console
func main() {
	ctx := context.Background()
	sub, err := pubsub.OpenSubscription(ctx, "gcppubsub://eliben-test1/test1")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := sub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got Message", msg)
	msg.Ack()
}
