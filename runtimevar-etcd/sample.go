package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
	"gocloud.dev/runtimevar"
	"gocloud.dev/runtimevar/etcdvar"
)

// MyConfig is a sample configuration struct.
type MyConfig struct {
	Server string
	Port   int
}

func main() {
	// Connect to the etcd server.
	client, err := clientv3.NewFromURL("localhost:2379")
	if err != nil {
		log.Fatal(err)
	}

	// Create a decoder for decoding JSON strings into MyConfig.
	decoder := runtimevar.NewDecoder(MyConfig{}, runtimevar.JSONDecode)

	// Construct a *runtimevar.Variable that watches the variable.
	// The etcd variable being referenced should have a JSON string that
	// decodes into MyConfig.
	v, err := etcdvar.OpenVariable(client, "cfg-variable-name", decoder, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer v.Close()

	// We can now read the current value of the variable from v.
	ctx, cancelFunc := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancelFunc()
	snapshot, err := v.Latest(ctx)
	if err != nil {
		log.Fatal(err)
	}
	cfg := snapshot.Value.(MyConfig)
	fmt.Println(cfg)
}
