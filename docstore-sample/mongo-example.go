// After starting up a local mongo server with the default settings/port, run:
// $ MONGO_SERVER_URL="mongodb://localhost" go run mongo-example.go
package main

import (
	"context"
	"log"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

type Player struct {
	Name             string
	Score            int
	DocstoreRevision interface{}
}

func main() {

	ctx := context.Background()

	// docstore.OpenCollection creates a *docstore.Collection from a URL.
	coll, err := docstore.OpenCollection(ctx, "mongo://my-db/my-collection")
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()
}
