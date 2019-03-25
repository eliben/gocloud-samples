package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/eliben/gocdkx/blob"
	_ "github.com/eliben/gocdkx/blob/gcsblob"
)

func main() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	iter := b.List(nil)
	for {
		obj, err := iter.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(obj.Key)
	}
}
