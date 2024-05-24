package containers

import (
	"time"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Log struct {
	Created time.Time `json:"created"`
	Docid   string    `json:"docid"`
	Stream  string    `json:"stream"`
	Text    string    `json:"text"`
}
type ContainerLogsResponse struct {
	api.ListResponse
	Logs []Log `json:"logs"`
}

type NetworkSettings struct {
	Networks struct {
		Bridge struct {
			Aliases             interface{} `json:"Aliases"`
			DriverOpts          interface{} `json:"DriverOpts"`
			EndpointID          string      `json:"EndpointID"`
			Gateway             string      `json:"Gateway"`
			GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
			GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
			IPAMConfig          interface{} `json:"IPAMConfig"`
			IPAddress           string      `json:"IPAddress"`
			IPPrefixLen         int         `json:"IPPrefixLen"`
			IPv6Gateway         string      `json:"IPv6Gateway"`
			Links               interface{} `json:"Links"`
			MacAddress          string      `json:"MacAddress"`
			NetworkID           string      `json:"NetworkID"`
		} `json:"bridge,omitempty"`
		Host struct {
			Aliases             interface{} `json:"Aliases"`
			DriverOpts          interface{} `json:"DriverOpts"`
			EndpointID          string      `json:"EndpointID"`
			Gateway             string      `json:"Gateway"`
			GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
			GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
			IPAMConfig          interface{} `json:"IPAMConfig"`
			IPAddress           string      `json:"IPAddress"`
			IPPrefixLen         int         `json:"IPPrefixLen"`
			IPv6Gateway         string      `json:"IPv6Gateway"`
			Links               interface{} `json:"Links"`
			MacAddress          string      `json:"MacAddress"`
			NetworkID           string      `json:"NetworkID"`
		} `json:"host,omitempty"`
	} `json:"Networks"`
}

type State struct {
	Dead       bool      `json:"Dead"`
	Error      string    `json:"Error"`
	ExitCode   int       `json:"ExitCode"`
	FinishedAt time.Time `json:"FinishedAt"`
	FinishedTS int64     `json:"FinishedTs"`
	OOMKilled  bool      `json:"OOMKilled"`
	Paused     bool      `json:"Paused"`
	Pid        int       `json:"Pid"`
	Restarting bool      `json:"Restarting"`
	Running    bool      `json:"Running"`
	StartedAt  time.Time `json:"StartedAt"`
	StartedTS  int       `json:"StartedTs"`
	Status     string    `json:"Status"`
}
type Container struct {
	Labels              map[string]string `json:"Labels"`
	NetworkSettings     *NetworkSettings  `json:"NetworkSettings"`
	State               *State            `json:"State"`
	Cmd                 string            `json:"cmd"`
	Created             int               `json:"created"`
	EnableServicePortal bool              `json:"enable_service_portal"`
	Exporting           bool              `json:"exporting"`
	FinishTime          int64             `json:"finish_time"`
	ID                  string            `json:"id"`
	Image               string            `json:"image"`
	IsDdsm              bool              `json:"is_ddsm"`
	IsPackage           bool              `json:"is_package"`
	Name                string            `json:"name"`
	Services            interface{}       `json:"services"`
	Status              string            `json:"status"`
	UpStatus            string            `json:"up_status"`
	UpTime              int               `json:"up_time"`
}

type ContainerList struct {
	Containers []Container `json:"containers"`
	api.ListResponse
}

type ContainerRestart struct {
	CPU           float64 `json:"cpu"`
	Memory        int     `json:"memory"`
	MemoryPercent float64 `json:"memoryPercent"`
	Name          string  `json:"name"`
}
