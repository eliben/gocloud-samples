package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eliben/gocdkx/contrib/secrets/vault"
	"github.com/hashicorp/vault/api"
)

// To run this, first start a Vault server running.
//
// An easy way is to run ./secrets/vault/localvault.sh from the root dir of
// the go-cloud repository (it sets VAULT_DEV_ROOT_TOKEN_ID to "faketoken").
func main() {
	ctx := context.Background()
	client, err := vault.Dial(ctx, &vault.Config{
		Token: "faketoken",
		APIConfig: api.Config{
			Address: "http://127.0.0.1:8200",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if _, err := client.Logical().Write("sys/mounts/transit", map[string]interface{}{"type": "transit"}); err != nil {
		fmt.Println("Error while mounting\n----", err)
		fmt.Println("----")
	}

	// Construct a *secrets.Keeper.
	keeper := vault.OpenKeeper(client, "my-key", nil)
	defer keeper.Close()

	// Now we can use keeper to encrypt or decrypt.
	plaintext := []byte("Hello, Secrets!")
	ciphertext, err := keeper.Encrypt(ctx, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	decrypted, err := keeper.Decrypt(ctx, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(decrypted)
}
