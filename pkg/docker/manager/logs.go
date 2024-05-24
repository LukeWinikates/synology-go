package manager

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

func (c *client) GetContainerManagerLogs() (*api.ResponseWrapper[*ContainerManagerLogs], error) {
	return api.PerformRequest[*ContainerManagerLogs](c.apiClient, func(query url.Values) {
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
}
