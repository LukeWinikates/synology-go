package docker

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

func (c *client) GetContainer(name string) (string, error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "get")
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

func (c *client) GetContainerManagerLogs() (*api.ResponseWrapper[*ContainerManagerLogs], error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Log")
		query.Add("version", "1")
		query.Add("method", "list")
		query.Add("action", `"load"'`)
		query.Add("offset", "0")
		query.Add("limit", "100")

		query.Add("datefrom", `0`)
		query.Add("dateto", `0`)
		query.Add("loglevel", `""`)
		query.Add("filter_content", `""`)
		query.Add("sort_by", `"time"`)
		query.Add("sort_dir", `"ASC"`)
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return api.ParseResponse[*ContainerManagerLogs](resp.Body)
}

func (c *client) GetContainerLogs(name string) (*api.ResponseWrapper[*ContainerLogsResponse], error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container.Log")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("name", fmt.Sprintf("\"%s\"", name))
		query.Add("offset", "0")
		query.Add("limit", "100")

		query.Add("from", `""`)
		query.Add("to", `""`)
		query.Add("level", `""`)
		query.Add("keyword", "")
		query.Add("sort_by", "time")
		query.Add("sort_dir", `"ASC"`)
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return api.ParseResponse[*ContainerLogsResponse](resp.Body)
}
func (c *client) ListContainers() (*api.ResponseWrapper[*ContainerList], error) {
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "list")
		query.Add("limit", "-1")
		query.Add("offset", "0")
		query.Add("type", "all")
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return api.ParseResponse[*ContainerList](resp.Body)
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
	req, err := c.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Container")
		query.Add("version", "1")
		query.Add("method", "restart")
		query.Add("name", name)
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return api.ParseResponse[*ContainerRestart](resp.Body)
}
