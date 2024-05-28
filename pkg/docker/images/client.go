package images

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Client interface {
	Get(imageName, tag string) (*api.ResponseWrapper[*Image], error)
	StartUpgradeCheck(repository string) (*api.ResponseWrapper[*Task], error)
	GetUpgradeTaskStatus(taskID string) (*api.ResponseWrapper[*UpgradeStatus], error)
	List() (*api.ResponseWrapper[*ImageList], error)
	StartPull(repository string, tag string) (*api.ResponseWrapper[*Task], error)
	GetPullStatus(taskID string) (*api.ResponseWrapper[*PullStatus], error)
}

type client struct {
	apiClient api.Client
}

func (c *client) StartPull(repository string, tag string) (*api.ResponseWrapper[*Task], error) {
	return api.PerformRequest[*Task](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "pull_start")
		query.Add("repository", repository)
		query.Add("tag", tag)

	})
}

func (c *client) GetPullStatus(taskID string) (*api.ResponseWrapper[*PullStatus], error) {
	return api.PerformRequest[*PullStatus](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "pull_status")
		query.Add("task_id", api.WrapQuote(taskID))
	})

}

func (c *client) Get(imageName, tag string) (*api.ResponseWrapper[*Image], error) {
	return api.PerformRequest[*Image](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("image", imageName)
		query.Add("tag", tag)
	})
}

func (c *client) StartUpgradeCheck(repository string) (*api.ResponseWrapper[*Task], error) {
	return api.PerformRequest[*Task](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "upgrade_start")
		query.Add("repository", repository)
	})
}

func (c *client) GetUpgradeTaskStatus(taskID string) (*api.ResponseWrapper[*UpgradeStatus], error) {
	return api.PerformRequest[*UpgradeStatus](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "upgrade_status")
		query.Add("task_id", api.WrapQuote(taskID))
	})
}

func (c *client) List() (*api.ResponseWrapper[*ImageList], error) {
	return api.PerformRequest[*ImageList](c.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Image")
		query.Add("version", "1")
		query.Add("method", "list")
		query.Add("limit", "-1")
		query.Add("offset", "0")
		query.Add("show_dsm", "false")
	})
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}
