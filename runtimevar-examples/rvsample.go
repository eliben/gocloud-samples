package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/runtimevar"
	_ "gocloud.dev/runtimevar/runtimeconfigurator"
	runtimeconfig "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
	"google.golang.org/grpc/status"
)

func snpsht() {
	ctx := context.Background()
	v, err := runtimevar.OpenVariable(ctx, "runtimeconfigurator://eliben-test1/eliben-testconfig/key1")
	if err != nil {
		log.Fatal(err)
	}

	s, err := v.Latest(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s.Value)

	var rcv *runtimeconfig.Variable
	if s.As(&rcv) {
		fmt.Println(rcv.UpdateTime)
	}
}

func erroras() {
	ctx := context.Background()
	v, err := runtimevar.OpenVariable(ctx, "runtimeconfigurator://eliben-test1/eliben-teconfig/key1g")
	if err != nil {
		log.Fatal(err)
	}

	_, err = v.Watch(ctx)
	if err != nil {
		var s *status.Status
		if v.ErrorAs(err, &s) {
			fmt.Println(s.Code())
		}
	}
}

func main() {
	snpsht()
	erroras()
}
