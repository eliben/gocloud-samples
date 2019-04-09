package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/runtimevar"
	_ "gocloud.dev/runtimevar/gcpruntimeconfig"
)

func main() {
	ctx := context.Background()
	v, err := runtimevar.OpenVariable(ctx, "gcpruntimeconfig://eliben-test1/eliben-testconfig/key1")
	if err != nil {
		log.Fatal(err)
	}
	s, err := v.Latest(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Value)
}
