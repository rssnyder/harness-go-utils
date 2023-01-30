package secrets

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/harness/harness-go-sdk/harness/nextgen"

	"github.com/rssnyder/harness-go-utils/config"
)

// SetSecretText creates/updates a text secret
func SetSecretText(ctx context.Context, client *nextgen.APIClient, identifier, name, content, secretManager string) (err error) {
	if secretManager == "" {
		secretManager = "harnessSecretManager"
	}

	organization, project := config.GetScope()

	secret := &nextgen.Secret{
		Type_:      nextgen.SecretTypes.SecretText,
		Name:       name,
		Identifier: identifier,
		Text: &nextgen.SecretTextSpec{
			Type_:                   nextgen.SecretSpecTypes.Text,
			ValueType:               nextgen.SecretTextValueTypes.Inline,
			Value:                   content,
			SecretManagerIdentifier: "harnessSecretManager",
		},
	}
	if organization.IsSet() {
		secret.OrgIdentifier = organization.Value()
	}
	if project.IsSet() {
		secret.ProjectIdentifier = project.Value()
	}

	resp, _, err := client.SecretsApi.GetSecretV2(ctx, identifier, client.AccountId, &nextgen.SecretsApiGetSecretV2Opts{
		OrgIdentifier:     organization,
		ProjectIdentifier: project,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	if resp.Data == nil {
		_, _, err = client.SecretsApi.PostSecret(ctx, nextgen.SecretRequestWrapper{Secret: secret}, client.AccountId, &nextgen.SecretsApiPostSecretOpts{
			OrgIdentifier:     organization,
			ProjectIdentifier: project,
		})
		fmt.Println(err)
		return err
	} else {
		_, _, err = client.SecretsApi.PutSecret(ctx, client.AccountId, identifier, &nextgen.SecretsApiPutSecretOpts{
			Body:              optional.NewInterface(nextgen.SecretRequestWrapper{Secret: secret}),
			OrgIdentifier:     organization,
			ProjectIdentifier: project,
		})
		fmt.Println(err)
		return err
	}
}
