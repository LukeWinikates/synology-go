package containers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

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
