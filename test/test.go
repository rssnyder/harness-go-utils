package main

import (
	"fmt"

	"github.com/rssnyder/harness-go-utils/config"
	"github.com/rssnyder/harness-go-utils/secrets"
)

func main() {
	c, ctx := config.GetNextgenClient()
	err := secrets.SetSecretText(ctx, c, "cli", "cli", "cli", "")
	if err != nil {
		fmt.Println(err)
	}
	err = secrets.SetSecretText(ctx, c, "cli0", "cli", "cli", "")
	if err != nil {
		fmt.Println(err)
	}
}
