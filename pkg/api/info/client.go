package info

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Info map[string]Description
type Description struct {
	MaxVersion    int    `json:"maxVersion"`
	MinVersion    int    `json:"minVersion"`
	Path          string `json:"path"`
	RequestFormat string `json:"requestFormat"`
}

type Client interface {
	GetInfo() (*api.ResponseWrapper[*Info], error)
}

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{apiClient: apiClient}
}

func (c *client) GetInfo() (*api.ResponseWrapper[*Info], error) {
	return api.GET[*Info](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.API.Info")
		query.Add("version", "1")
		query.Add("method", "query")
	})
}
