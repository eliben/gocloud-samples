package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/gcsblob"
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
