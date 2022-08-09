package mailslurp

import (
	"context"
	mailslurpapi "github.com/mailslurp/mailslurp-client-go"
)

// Config struct holds API key
//
type Config struct {
	MailSlurpClient *mailslurpapi.APIClient
	Ctx             context.Context
}

func ConfigureClient(apiKey string) (Config, error) {
	ctx := context.WithValue(context.Background(), mailslurpapi.ContextAPIKey, mailslurpapi.APIKey{Key: apiKey})
	config := mailslurpapi.NewConfiguration()
	return Config{
		MailSlurpClient: mailslurpapi.NewAPIClient(config),
		Ctx:             ctx,
	}, nil
}
