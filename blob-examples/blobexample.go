package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/storage"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"gocloud.dev/blob"
	"gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
	"gocloud.dev/gcp"
	"google.golang.org/api/googleapi"
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

func errortype() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	// Try read a file that doesn't exist
	_, err = b.ReadAll(ctx, "XOXOXO")
	if err != nil {
		log.Printf("ReadAll %v %T\n", err, err)

		var gError *googleapi.Error
		if b.ErrorAs(err, &gError) {
			log.Printf("Converted to %T: %v\n", *gError, *gError)
		} else {
			log.Printf("Failed to convert to specific error\n")
		}
	}
}

func list() {
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
		var oa storage.ObjectAttrs
		if obj.As(&oa) {
			fmt.Println(oa.Owner)
		}
	}
}

func listopt() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	beforeList := func(as func(interface{}) bool) error {
		var q *storage.Query
		if as(&q) {
			fmt.Println(q.Delimiter)
		}
		return nil
	}

	iter := b.List(&blob.ListOptions{Prefix: "", Delimiter: "/", BeforeList: beforeList})
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

func reader() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	r, err := b.NewReader(ctx, "gopher.png", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var sr storage.Reader
	if r.As(&sr) {
		fmt.Println(sr.Attrs)
	}
}

func attrs() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	attrs, err := b.Attributes(ctx, "gopher.png")
	if err != nil {
		log.Fatal(err)
	}

	var oa storage.ObjectAttrs
	if attrs.As(&oa) {
		fmt.Println(oa.Owner)
	}
}

func writeopt() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "gs://eliben-test-bucket")
	if err != nil {
		log.Fatal(err)
	}

	beforeWrite := func(as func(interface{}) bool) error {
		var sw *storage.Writer
		if as(&sw) {
			fmt.Println(sw.ChunkSize)
		}
		return nil
	}

	options := blob.WriterOptions{BeforeWrite: beforeWrite}
	if err := b.WriteAll(ctx, "newfile.txt", []byte("hello\n"), &options); err != nil {
		log.Fatal(err)
	}
}

func aserror() {
	ctx := context.Background()

	b, err := blob.OpenBucket(ctx, "s3://eliben-testing")
	if err != nil {
		log.Fatal(err)
	}

	_, err = b.ReadAll(ctx, "nosuchfile")
	if err != nil {
		var awsErr awserr.Error
		if b.ErrorAs(err, &awsErr) {
			fmt.Println(awsErr.Code())
		}
	}
}

func main() {
	//full()
	//url()
	//errortype()
	//list()
	//listopt()
	//reader()
	//attrs()
	//writeopt()
	aserror()
}
