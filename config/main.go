package config

import (
	"context"
	"os"
	"strings"
	"sync"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"
)

var (
	configureClient sync.Once
	client          *nextgen.APIClient
)

// getNextgenClient configures a client and context using env vars
// HARNESS_API_KEY: harness nextgen api key
// HARNESS_ACCOUNT_ID: harness account id
func GetNextgenClient() (client *nextgen.APIClient, ctx context.Context) {
	configureClient.Do(func() {
		cfg := nextgen.NewConfiguration()
		client = nextgen.NewAPIClient(cfg)
	})

	ctx = context.WithValue(context.Background(), nextgen.ContextAPIKey, nextgen.APIKey{Key: client.ApiKey})

	return
}

// getOrgProj pulls org and project information using env vars, specifying neither results in account level scoping
// HARNESS_PLATFORM_ORGANIZATION: organization id
// HARNESS_PLATFORM_PROJECT: project id
func GetScope() (organization, project optional.String) {
	if value, ok := os.LookupEnv("HARNESS_PLATFORM_ORGANIZATION"); ok {
		organization = optional.NewString(value)
	} else {
		organization = optional.EmptyString()
	}

	if value, ok := os.LookupEnv("HARNESS_PLATFORM_PROJECT"); ok {
		project = optional.NewString(value)
	} else {
		project = optional.EmptyString()
	}

	return
}

// NameToIdentifier formats names into harness ids
func NameToIdentifier(name string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(name, " ", "_"), "-", ""))
}
