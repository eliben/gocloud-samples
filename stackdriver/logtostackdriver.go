package main

import (
	"log"

	// Imports the Stackdriver Logging client package.
	"cloud.google.com/go/logging"
	"golang.org/x/net/context"
)

// To see these in the stackdriver console I had to select "Project" rather
// than "Global" from the filter. Discovered this by running:
// $ gcloud logging read my-test-log
func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "eliben-test1"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the log to write to.
	//logName := "my-test-log"

	//logger := client.Logger(logName).StandardLogger(logging.Info)
	logger := client.Logger("my-test-log")

	// Logs "hello world", log entry is visible at
	// Stackdriver Logs.
	//logger.Println("hello world 1221")
	logger.LogSync(ctx, logging.Entry{Payload: "sync"})

	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}
}
