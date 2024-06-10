package containers

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
)

// Client calls DSM APIs such as "SYNO.Docker.Container" to query or modify Container Manager containers
type Client interface {
	ListContainers() (*api.ResponseWrapper[*ContainerList], error)
	RestartContainer(name string) (*api.ResponseWrapper[*ContainerStats], error)
	StopContainer(name string) (*api.ResponseWrapper[*ContainerStats], error)
	StartContainer(name string) (*api.ResponseWrapper[*ContainerStats], error)
	GetContainer(name string) (*api.ResponseWrapper[*DetailsAndProfile], error)
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
