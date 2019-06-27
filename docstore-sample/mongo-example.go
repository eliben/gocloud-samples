// After starting up a local mongo server with the default settings/port, run:
// $ MONGO_SERVER_URL="mongodb://localhost" go run mongo-example.go
package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/mongodocstore"
	"gocloud.dev/gcerrors"
)

type Player struct {
	Name             string
	Score            int
	DocstoreRevision interface{}
}

func main() {
	ctx := context.Background()

	coll, err := docstore.OpenCollection(ctx, "mongo://my-db/my-collection?id_field=Name")
	if err != nil {
		log.Fatal(err)
	}
	defer coll.Close()

	name := "Patr"
	fmt.Println("Fetching player", name)
	p := &Player{Name: name}
	err = coll.Get(ctx, p)
	if err != nil {
		// if not found, insert
		if gcerrors.Code(err) == gcerrors.NotFound {
			fmt.Println("Inserting")
			err = coll.Create(ctx, &Player{Name: name, Score: 10})
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Fatal(err)
	}
	fmt.Println(p)
}
