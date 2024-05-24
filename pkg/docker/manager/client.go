package manager

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
)

// Client calls DSM APIs such as "SYNO.Docker.Log" and "SYNO.Docker.Container" to query or modify Container Manager containers
type Client interface {
	GetContainerManagerLogs() (*api.ResponseWrapper[*ContainerManagerLogs], error)
}

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}
