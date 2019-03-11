package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/runtimevar"
	_ "gocloud.dev/runtimevar/runtimeconfigurator"
	runtimeconfig "google.golang.org/genproto/googleapis/cloud/runtimeconfig/v1beta1"
)

func main() {
	ctx := context.Background()

	v, err := runtimevar.OpenVariable(ctx, "runtimeconfigurator://eliben-test1/eliben-testconfig/key1?decoder=string")
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
