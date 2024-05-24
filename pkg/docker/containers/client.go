package containers

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
)

// Client calls DSM APIs such as "SYNO.Docker.Container" to query or modify Container Manager containers
type Client interface {
	ListContainers() (*api.ResponseWrapper[*ContainerList], error)
	RestartContainer(name string) (*api.ResponseWrapper[*ContainerRestart], error)
	StopContainer(name string) (string, error)
	StartContainer(name string) (string, error)
	GetContainer(name string) (string, error)
	GetContainerLogs(name string) (*api.ResponseWrapper[*ContainerLogsResponse], error)
}

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}
