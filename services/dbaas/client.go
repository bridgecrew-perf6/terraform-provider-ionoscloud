package dbaas

import (
	dbaas "github.com/ionos-cloud/sdk-go-dbaas-postgres"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/utils"
	"net/http"
	"os"
)

type Client struct {
	dbaas.APIClient
}

type ClientConfig struct {
	dbaas.Configuration
}

// ClientService is a wrapper around dbaas.APIClient
type ClientService interface {
	Get() *Client
	GetConfig() *ClientConfig
}

type clientService struct {
	client *dbaas.APIClient
}

//var _ ClientService = &clientService{}

func NewClientService(username, password, token, url string) ClientService {
	newConfigDbaas := dbaas.NewConfiguration(username, password, token, url)

	if os.Getenv(utils.IonosDebug) != "" {
		newConfigDbaas.Debug = true
	}
	newConfigDbaas.MaxRetries = utils.MaxRetries
	newConfigDbaas.MaxWaitTime = utils.MaxWaitTime

	newConfigDbaas.HTTPClient = &http.Client{Transport: utils.CreateTransport()}

	return &clientService{
		client: dbaas.NewAPIClient(newConfigDbaas),
	}
}

func (c clientService) Get() *Client {
	return &Client{
		APIClient: *c.client,
	}
}

func (c clientService) GetConfig() *ClientConfig {
	return &ClientConfig{
		Configuration: *c.client.GetConfig(),
	}
}
