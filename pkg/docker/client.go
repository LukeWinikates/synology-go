package docker

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Client interface {
	ListContainers() (*api.ResponseWrapper[*ContainerList], error)
	RestartContainer(name string) (*api.ResponseWrapper[*ContainerRestart], error)
	StopContainer(name string) (string, error)
	StartContainer(name string) (string, error)
	GetContainer(name string) (string, error)
}

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}
