package main

import (
	"context"
	"fmt"
	"log"

	_ "github.com/eliben/gocloudsampleprovider"
	"gocloud.dev/blob"
)

func main() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gcspfile://tempfile1")
	if err != nil {
		log.Fatal(err)
	}
	err = b.WriteAll(ctx, "my-key", []byte("hello world"), nil)
	if err != nil {
		log.Fatal(err)
	}
	data, err := b.ReadAll(ctx, "my-key")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
