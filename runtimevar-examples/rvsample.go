package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/runtimevar"
	_ "gocloud.dev/runtimevar/runtimeconfigurator"
)

func main() {
	ctx := context.Background()

	v, err := runtimevar.OpenVariable(ctx, "runtimeconfigurator://eliben-test1/eliben-testconfig/key1?decoder=string")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
