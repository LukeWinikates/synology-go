package containers

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

func (c *client) GetContainer(name string) (*api.ResponseWrapper[*DetailsAndProfile], error) {
	return api.PerformRequest[*DetailsAndProfile](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("name", name)
	})
}

func (c *client) ListContainers() (*api.ResponseWrapper[*ContainerList], error) {
	return api.PerformRequest[*ContainerList](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "list")
		query.Add("limit", "-1")
		query.Add("offset", "0")
		query.Add("type", "all")
	})
}

func (c *client) StopContainer(name string) (*api.ResponseWrapper[*ContainerStats], error) {
	return api.PerformRequest[*ContainerStats](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "stop")
		query.Add("name", name)
	})
}

func (c *client) StartContainer(name string) (*api.ResponseWrapper[*ContainerStats], error) {
	return api.PerformRequest[*ContainerStats](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "start")
		query.Add("name", name)
	})
}

func (c *client) RestartContainer(name string) (*api.ResponseWrapper[*ContainerStats], error) {
	return api.PerformRequest[*ContainerStats](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "restart")
		query.Add("name", name)
	})
}
