package main

import (
	"context"
	"fmt"
	"log"

	"gocloud.dev/secrets"
	_ "gocloud.dev/secrets/gcpkms"
	"google.golang.org/grpc/status"
)

func do() {
	ctx := context.Background()

	k, err := secrets.OpenKeeper(ctx, "gcpkms://projects/eliben-test1/locations/global/keyRings/test/cryptoKeys/quickstart")
	if err != nil {
		log.Fatal(err)
	}

	plaintext := []byte("Go CDK secrets")
	ciphertext, err := k.Encrypt(ctx, plaintext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ciphertext)
}

func erroras() {
	ctx := context.Background()

	k, err := secrets.OpenKeeper(ctx, "gcpkms://projects/eliben-test1/locations/global/keyRings/test/wrong/quickstart")
	if err != nil {
		log.Fatal("open", err)
	}

	plaintext := []byte("Go CDK secrets")
	ciphertext, err := k.Encrypt(ctx, plaintext)
	if err != nil {
		var s *status.Status
		if k.ErrorAs(err, &s) {
			fmt.Println(s.Code())
		}
	}

	fmt.Println(ciphertext)
}

func main() {
	do()
	erroras()
}
