package manager

import "github.com/LukeWinikates/synology-go/pkg/api"

type ContainerManagerLog struct {
	Event   string `json:"event"`
	Level   string `json:"level"`
	LogType string `json:"log_type"`
	Time    string `json:"time"`
	User    string `json:"user"`
}

type ContainerManagerLogs struct {
	api.ListResponse
	Logs       []ContainerManagerLog `json:"logs"`
	ErrorCount int                   `json:"error_count"`
	InfoCount  int                   `json:"info_count"`
	WarnCount  int                   `json:"warn_count"`
}
