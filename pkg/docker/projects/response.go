package projects

import "time"

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

type DeleteResponse struct{}
