package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
)

func main() {
	ctx := context.Background()
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Create an HTTP client.
	// This example uses the default HTTP transport and the credentials created
	// above.
	client, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
	if err != nil {
		return
	}

	// Create a *blob.Bucket.
	b, err := gcsblob.OpenBucket(ctx, client, "eliben-test-bucket", nil)
	if err != nil {
		log.Fatal("OpenBucket", err)
	}

	// Now we can use b to read or write files to the container.
	data, err := b.ReadAll(ctx, "gopher.png")
	if err != nil {
		log.Fatal("ReadAll", err)
	}

	fmt.Println(data)
}
