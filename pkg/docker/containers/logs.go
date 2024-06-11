package containers

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

const APISynoDockerContainerLog = "SYNO.Docker.Container.Log"

func (c *client) GetContainerLogs(name string) (*api.ResponseWrapper[*ContainerLogsResponse], error) {
	return api.PerformRequest[*ContainerLogsResponse](c.apiClient, func(query url.Values) {
		query.Add("api", APISynoDockerContainerLog)
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("name", api.WrapQuote(name))
		query.Add("offset", "0")
		query.Add("limit", "100")

		query.Add("from", `""`)
		query.Add("to", `""`)
		query.Add("level", `""`)
		query.Add("keyword", "")
		query.Add("sort_by", "time")
		query.Add("sort_dir", `"ASC"`)
	})
}
