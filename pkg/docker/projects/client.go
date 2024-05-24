package projects

import (
	"fmt"
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type client struct {
	apiClient api.Client
}

func NewClient(apiClient api.Client) Client {
	return &client{
		apiClient: apiClient,
	}
}

type Client interface {
	List() (*api.ResponseWrapper[ProjectList], error)
	Get(id string) (*api.ResponseWrapper[*Project], error)
	BuildStream(id string, lineReader func(s string)) error
	Create(name string, sharePath string, content string, servicePortalConfig *ServicePortalConfiguration) (*api.ResponseWrapper[*Project], error)
	GetShareInfo(path string) (*api.ResponseWrapper[*ShareInfo], error)
	Delete(id string) (*api.ResponseWrapper[*DeleteResponse], error)
	UpdateContent(id string, content string) (*api.ResponseWrapper[*Project], error)
	Stop(id string, lineReader func(s string)) error
	Start(id string, lineReader func(s string)) error
}

func (pc *client) GetShareInfo(path string) (*api.ResponseWrapper[*ShareInfo], error) {
	return api.PerformRequest[*ShareInfo](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "get_share_info")
		query.Add("path", fmt.Sprintf(`"%s"`, path))
	})
}

func (pc *client) Get(id string) (*api.ResponseWrapper[*Project], error) {
	return api.PerformRequest[*Project](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("id", fmt.Sprintf(`"%s"`, id))
	})
}

func (pc *client) Create(name string, sharePath, content string, servicePortalConfig *ServicePortalConfiguration) (*api.ResponseWrapper[*Project], error) {
	return api.PerformRequest[*Project](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "create")

		query.Add("name", fmt.Sprintf(`"%s"`, name))
		query.Add("content", fmt.Sprintf(`"%s"`, content))
		query.Add("share_path", sharePath)

		applyServicePortalOptions(servicePortalConfig, query)
	})
}

func (pc *client) Delete(id string) (*api.ResponseWrapper[*DeleteResponse], error) {
	return api.PerformRequest[*DeleteResponse](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "delete")
		query.Add("id", id)
	})
}

func (pc *client) List() (*api.ResponseWrapper[ProjectList], error) {
	return api.PerformRequest[ProjectList](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "list")
	})
}

func (pc *client) UpdateContent(id, content string) (*api.ResponseWrapper[*Project], error) {
	return api.PerformPOSTRequest[*Project](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "update")
		query.Add("id", fmt.Sprintf(`"%s"`, id))
		query.Add("content", content)
	})
}
