package docker

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Project struct {
	ContainerIDs          []string    `json:"containerIds"`
	CreatedAt             time.Time   `json:"created_at"`
	EnableServicePortal   bool        `json:"enable_service_portal"`
	ID                    string      `json:"id"`
	IsPackage             bool        `json:"is_package"`
	Name                  string      `json:"name"`
	Path                  string      `json:"path"`
	ServicePortalName     string      `json:"service_portal_name"`
	ServicePortalPort     int         `json:"service_portal_port"`
	ServicePortalProtocol string      `json:"service_portal_protocol"`
	Services              interface{} `json:"services"`
	SharePath             string      `json:"share_path"`
	State                 string      `json:"state"`
	Status                string      `json:"status"`
	UpdatedAt             time.Time   `json:"updated_at"`
	Version               int         `json:"version"`
}

type ProjectList = map[string]Project

type ShareInfo struct {
	ComposePath             string `json:"compose_path"`
	Content                 string `json:"content"`
	IsDockerComposeYmlExist bool   `json:"is_docker_compose_yml_exist"`
}
type projectClient struct {
	apiClient api.Client
}

func (pc *projectClient) GetShareInfo(path string) (*api.ResponseWrapper[*ShareInfo], error) {
	return api.PerformRequest[*ShareInfo](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "get_share_info")
		query.Add("path", fmt.Sprintf(`"%s"`, path))
	})
}

func (pc *projectClient) Get(id string) (*api.ResponseWrapper[*Project], error) {
	return api.PerformRequest[*Project](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "get")
		query.Add("id", id)
	})
}

func (pc *projectClient) BuildStream(id string, lineReader func(s string)) error {
	req, err := pc.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "build_stream")
		query.Add("id", id)
	})
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lineReader(scanner.Text())
	}
	return err
}

type ServicePortalConfiguration struct {
	Enabled  bool
	Name     string
	Port     int
	Protocol string
}

type DeleteResponse struct{}

func (pc *projectClient) Create(name string, sharePath, content string, servicePortalConfig *ServicePortalConfiguration) (*api.ResponseWrapper[*Project], error) {
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

func applyServicePortalOptions(_ *ServicePortalConfiguration, query url.Values) {
	query.Add("enable_service_portal", "false")
	query.Add("service_portal_name", `""`)
	query.Add("service_portal_port", "0")
	query.Add("service_portal_protocol", `"0"`)
}

func (pc *projectClient) Delete(id string) (*api.ResponseWrapper[*DeleteResponse], error) {
	return api.PerformRequest[*DeleteResponse](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "delete")
		query.Add("id", id)
	})
}

func (pc *projectClient) List() (*api.ResponseWrapper[ProjectList], error) {
	return api.PerformRequest[ProjectList](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "list")
	})
}

func (pc *projectClient) UpdateContent(id, content string) (*api.ResponseWrapper[*Project], error) {
	return api.PerformRequest[*Project](pc.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", "update")
		query.Add("id", id)
		query.Add("content", fmt.Sprintf(`"%s"`, content))
	})
}

func NewProjectClient(apiClient api.Client) ProjectClient {
	return &projectClient{
		apiClient: apiClient,
	}
}

type ProjectClient interface {
	List() (*api.ResponseWrapper[ProjectList], error)
	Get(id string) (*api.ResponseWrapper[*Project], error)
	BuildStream(id string, lineReader func(s string)) error
	Create(name string, sharePath string, content string, servicePortalConfig *ServicePortalConfiguration) (*api.ResponseWrapper[*Project], error)
	GetShareInfo(path string) (*api.ResponseWrapper[*ShareInfo], error)
	Delete(id string) (*api.ResponseWrapper[*DeleteResponse], error)
}
