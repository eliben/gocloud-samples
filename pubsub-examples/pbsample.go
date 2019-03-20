package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/pubsub"
	_ "gocloud.dev/pubsub/gcppubsub"

	pbraw "cloud.google.com/go/pubsub/apiv1"
	pbapi "google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc/status"
)

func rcv() {
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

func msgas() {
	ctx := context.Background()
	sub, err := pubsub.OpenSubscription(ctx, "gcppubsub://eliben-test1/test1")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := sub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var pm *pbapi.PubsubMessage
	if msg.As(&pm) {
		fmt.Println(pm.GetAttributes())
	}
	fmt.Println("got Message", msg)
	msg.Ack()
}

func subas() {
	ctx := context.Background()
	sub, err := pubsub.OpenSubscription(ctx, "gcppubsub://eliben-test1/test1")
	if err != nil {
		log.Fatal(err)
	}
	var sc *pbraw.SubscriberClient
	if sub.As(&sc) {
		fmt.Println(sc.CallOptions)
	}
}

func suberroras() {
	ctx := context.Background()
	sub, err := pubsub.OpenSubscription(ctx, "gcppubsub://eliben-test1/test100")
	if err != nil {
		log.Fatal(err)
	}

	msg, err := sub.Receive(ctx)
	if err != nil {
		var s *status.Status
		if sub.ErrorAs(err, &s) {
			fmt.Println(s.Code())
		}
		log.Fatal(err)
	}
	msg.Ack()
}

func send() {
	ctx := context.Background()
	topic, err := pubsub.OpenTopic(ctx, "gcppubsub://eliben-test1/test1")
	if err != nil {
		log.Fatal(err)
	}

	err = topic.Send(ctx, &pubsub.Message{Body: []byte("hello")})
	if err != nil {
		log.Fatal(err)
	}
}

func topicas() {
	ctx := context.Background()
	topic, err := pubsub.OpenTopic(ctx, "gcppubsub://eliben-test1/test1")
	if err != nil {
		log.Fatal(err)
	}

	var pc *pbraw.PublisherClient
	if topic.As(&pc) {
		fmt.Println("converted")
	}
}

func topicerroras() {
	ctx := context.Background()
	topic, err := pubsub.OpenTopic(ctx, "gcppubsub://eliben-test1/test100")
	if err != nil {
		log.Fatal(err)
	}

	err = topic.Send(ctx, &pubsub.Message{Body: []byte("hello")})
	if err != nil {
		var s *status.Status
		if topic.ErrorAs(err, &s) {
			fmt.Println(s.Code())
		}
		log.Fatal(err)
	}
}

// To publish a message here, post a message to the subscription via the GCP
// console
func main() {
	//rcv()
	//msgas()
	//subas()
	//suberroras()
	//send()
	//topicas()
	topicerroras()
}
