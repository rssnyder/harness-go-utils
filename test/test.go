package main

import (
	"github.com/rssnyder/harness-go-utils/secrets"
)

func main() {
	secrets.SetSecretText("cli", "cli", "cli", "")
}
