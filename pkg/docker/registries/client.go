package registries

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}

type Registry struct {
	EnableRegistryMirror bool     `json:"enable_registry_mirror,omitempty"`
	EnableTrustSSC       bool     `json:"enable_trust_SSC"`
	MirrorUrls           []string `json:"mirror_urls,omitempty"`
	Name                 string   `json:"name"`
	Syno                 bool     `json:"syno"`
	URL                  string   `json:"url"`
	Username             string   `json:"username,omitempty"`
}

type RegistryList struct {
	api.ListResponse
	Registries []Registry `json:"registries"`
	Using      string     `json:"using"`
}

type Client interface {
	List() (*api.ResponseWrapper[RegistryList], error)
}

func (pc *client) List() (*api.ResponseWrapper[RegistryList], error) {
	return api.GET[RegistryList](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Registry")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("limit", "-1")
		query.Add("offset", "0")
	})
}
