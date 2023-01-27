package main

import (
	"github.com/rssnyder/harness-go-utils/config"
	"github.com/rssnyder/harness-go-utils/secrets"
)

func main() {
	c, ctx := config.GetNextgenClient()
	secrets.SetSecretText(ctx, c, "cli", "cli", "cli", "")
	secrets.SetSecretText(ctx, c, "cli0", "cli", "cli", "")
}
