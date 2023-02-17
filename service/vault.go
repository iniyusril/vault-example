package service

import (
	"context"
	"fmt"
	"os"

	vault "github.com/hashicorp/vault/api"
)

var res map[string]any

func InitConfiguration() error {
	config := vault.DefaultConfig()
	config.Address = os.Getenv("VAULT_ADDR")

	client, err := vault.NewClient(config)
	if err != nil {
		os.Exit(0)
	}

	client.SetToken(os.Getenv("VAULT_TOKEN"))

	ctx := context.Background()

	secret, err := client.KVv2("secret").Get(ctx, "kv")
	if err != nil {
		os.Exit(0)
	}

	

	res = secret.Data

	return nil

}

func GetConfiguration(key string) string {
	resGet := res[key]

	if resGet == nil {
		resGet = ""
	}

	result := fmt.Sprintf("%v", resGet)
	return result
}
