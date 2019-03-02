package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "s3://eliben-testing")
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

	bb, err := b.ReadAll(ctx, "dynarray.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(bb))
}
