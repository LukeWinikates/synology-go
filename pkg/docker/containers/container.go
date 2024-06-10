package containers

import (
	"io"
	"net/http"
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

func (c *client) GetContainer(name string) (*api.ResponseWrapper[*Container], error) {
	return api.PerformRequest[*Container](c.apiClient, func(query url.Values) {
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

func (c *client) StopContainer(name string) (string, error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "stop")
		query.Add("name", name)
	})
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	return string(b), err
}

func (c *client) StartContainer(name string) (string, error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "start")
		query.Add("name", name)
	})
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)

	return string(b), err
}

func (c *client) RestartContainer(name string) (*api.ResponseWrapper[*ContainerRestart], error) {
	return api.PerformRequest[*ContainerRestart](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "restart")
		query.Add("name", name)
	})
}
