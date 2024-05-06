package docker

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Methods:
//// create
//// delete
//// import
//// get
//// get_log
//// get_process
//// list
//// restart
//// set
//// set_start
//// signal
//// start
//// stop
//// stats
//// upload

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
	FinishedTs int64     `json:"FinishedTs"`
	OOMKilled  bool      `json:"OOMKilled"`
	Paused     bool      `json:"Paused"`
	Pid        int       `json:"Pid"`
	Restarting bool      `json:"Restarting"`
	Running    bool      `json:"Running"`
	StartedAt  time.Time `json:"StartedAt"`
	StartedTs  int       `json:"StartedTs"`
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
	Id                  string            `json:"id"`
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
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	Total      int         `json:"total"`
}

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
	b, err := io.ReadAll(resp.Body)

	return string(b), err
}
func (c *client) ListContainers() (*api.ResponseWrapper[*ContainerList], error) {
	// special thanks to: https://github.com/Xboarder56/SynoDockerContainerUpgrade/blob/dab41946149baf2f06dda8c77b9e909ccd502b31/update_containers.py#L88
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
	b, err := io.ReadAll(resp.Body)

	return string(b), err
}

type ContainerRestart struct {
	Cpu           float64 `json:"cpu"`
	Memory        int     `json:"memory"`
	MemoryPercent float64 `json:"memoryPercent"`
	Name          string  `json:"name"`
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
	return api.ParseResponse[*ContainerRestart](resp.Body)
}
