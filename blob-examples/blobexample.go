package main

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"gocloud.dev/blob"
	"gocloud.dev/blob/gcsblob"
	"gocloud.dev/gcp"
)

func full() {
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

	var gcsClient *storage.Client
	if b.As(&gcsClient) {
		email, err := gcsClient.ServiceAccount(ctx, "eliben-test1")
		if err != nil {
			log.Fatal("ServiceAccount", err)
		}
		log.Println("email", email)
	} else {
		log.Fatal("Unable to access storage.Client through Bucket.As")
	}

	// Now we can use b to read or write files to the container.
	data, err := b.ReadAll(ctx, "gopher.png")
	if err != nil {
		log.Fatal("ReadAll", err)
	}

	log.Println(len(data))
}

func url() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	var gcsClient *storage.Client
	if b.As(&gcsClient) {
		email, err := gcsClient.ServiceAccount(ctx, "eliben-test1")
		if err != nil {
			log.Fatal("ServiceAccount", err)
		}
		log.Println("email", email)
	} else {
		log.Fatal("Unable to access storage.Client through Bucket.As")
	}
}

func main() {
	full()
	url()
}
