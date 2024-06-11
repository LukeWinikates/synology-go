package containers

import (
	"time"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type Log struct {
	Created time.Time `json:"created"`
	DocID   string    `json:"docid"`
	Stream  string    `json:"stream"`
	Text    string    `json:"text"`
}
type ContainerLogsResponse struct {
	api.ListResponse
	Logs []Log `json:"logs"`
}

type Network struct {
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
}

type NetworkSettings struct {
	Bridge                 string             `json:"Bridge"`
	EndpointID             string             `json:"EndpointID"`
	Gateway                string             `json:"Gateway"`
	GlobalIPv6Address      string             `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen    int                `json:"GlobalIPv6PrefixLen"`
	HairpinMode            bool               `json:"HairpinMode"`
	IPAddress              string             `json:"IPAddress"`
	IPPrefixLen            int                `json:"IPPrefixLen"`
	IPv6Gateway            string             `json:"IPv6Gateway"`
	LinkLocalIPv6Address   string             `json:"LinkLocalIPv6Address"`
	LinkLocalIPv6PrefixLen int                `json:"LinkLocalIPv6PrefixLen"`
	MacAddress             string             `json:"MacAddress"`
	Ports                  interface{}        `json:"Ports"`
	SandboxID              string             `json:"SandboxID"`
	SandboxKey             string             `json:"SandboxKey"`
	SecondaryIPAddresses   interface{}        `json:"SecondaryIPAddresses"`
	SecondaryIPv6Addresses interface{}        `json:"SecondaryIPv6Addresses"`
	Networks               map[string]Network `json:"Networks"`
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

type ContainerStats struct {
	CPU           float64 `json:"cpu"`
	Memory        int     `json:"memory"`
	MemoryPercent float64 `json:"memoryPercent"`
	Name          string  `json:"name"`
}

type Mount struct {
	Destination string `json:"Destination"`
	Mode        string `json:"Mode"`
	Propagation string `json:"Propagation"`
	RW          bool   `json:"RW"`
	Source      string `json:"Source"`
	Type        string `json:"Type"`
}

type Config struct {
	AttachStderr bool                   `json:"AttachStderr"`
	AttachStdin  bool                   `json:"AttachStdin"`
	AttachStdout bool                   `json:"AttachStdout"`
	Cmd          interface{}            `json:"Cmd"`
	DDSM         bool                   `json:"DDSM"`
	Domainname   string                 `json:"Domainname"`
	Entrypoint   []string               `json:"Entrypoint"`
	Env          []string               `json:"Env"`
	ExposedPorts map[string]interface{} `json:"ExposedPorts"`
	Hostname     string                 `json:"Hostname"`
	Image        string                 `json:"Image"`
	Labels       map[string]string      `json:"Labels"`
	OnBuild      interface{}            `json:"OnBuild"`
	OpenStdin    bool                   `json:"OpenStdin"`
	StdinOnce    bool                   `json:"StdinOnce"`
	Tty          bool                   `json:"Tty"`
	User         string                 `json:"User"`
	Volumes      interface{}            `json:"Volumes"`
	WorkingDir   string                 `json:"WorkingDir"`
}

type EnvVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Profile struct {
	CapAdd                interface{}       `json:"CapAdd"`
	CapDrop               interface{}       `json:"CapDrop"`
	Cmd                   string            `json:"cmd"`
	CPUPriority           int               `json:"cpu_priority"`
	EnablePublishAllPorts bool              `json:"enable_publish_all_ports"`
	EnableRestartPolicy   bool              `json:"enable_restart_policy"`
	Enabled               bool              `json:"enabled"`
	EnvVariables          []EnvVariable     `json:"env_variables"`
	Exporting             bool              `json:"exporting"`
	ID                    string            `json:"id"`
	Image                 string            `json:"image"`
	IsDDSM                bool              `json:"is_ddsm"`
	IsPackage             bool              `json:"is_package"`
	Labels                map[string]string `json:"labels"`
	Links                 []interface{}     `json:"links"`
	MemoryLimit           int               `json:"memory_limit"`
	Name                  string            `json:"name"`
	Network               []struct {
		Driver string `json:"driver"`
		Name   string `json:"name"`
	} `json:"network"`
	NetworkMode  string `json:"network_mode"`
	PortBindings []struct {
		ContainerPort int    `json:"container_port"`
		HostPort      int    `json:"host_port"`
		Type          string `json:"type"`
	} `json:"port_bindings"`
	Privileged bool `json:"privileged"`
	Shortcut   struct {
		EnableShortcut   bool   `json:"enable_shortcut"`
		EnableStatusPage bool   `json:"enable_status_page"`
		EnableWebPage    bool   `json:"enable_web_page"`
		WebPageURL       string `json:"web_page_url"`
	} `json:"shortcut"`
	UseHostNetwork bool          `json:"use_host_network"`
	Version        int           `json:"version"`
	VolumeBindings []interface{} `json:"volume_bindings"`
}

type HostConfig struct {
	AutoRemove           bool          `json:"AutoRemove"`
	Binds                interface{}   `json:"Binds"`
	BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps"`
	BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps"`
	BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps"`
	BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps"`
	BlkioWeight          int           `json:"BlkioWeight"`
	BlkioWeightDevice    interface{}   `json:"BlkioWeightDevice"`
	CapAdd               interface{}   `json:"CapAdd"`
	CapDrop              interface{}   `json:"CapDrop"`
	Cgroup               string        `json:"Cgroup"`
	CgroupParent         string        `json:"CgroupParent"`
	CgroupnsMode         string        `json:"CgroupnsMode"`
	ConsoleSize          []int         `json:"ConsoleSize"`
	ContainerIDFile      string        `json:"ContainerIDFile"`
	CPUCount             int           `json:"CpuCount"`
	CPUPercent           int           `json:"CpuPercent"`
	CPUPeriod            int           `json:"CpuPeriod"`
	CPUQuota             int           `json:"CpuQuota"`
	CPURealtimePeriod    int           `json:"CpuRealtimePeriod"`
	CPURealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
	CPUShares            int           `json:"CpuShares"`
	CPUSetCPUs           string        `json:"CpusetCpus"`
	CPUSetMems           string        `json:"CpusetMems"`
	DeviceCgroupRules    interface{}   `json:"DeviceCgroupRules"`
	DeviceRequests       interface{}   `json:"DeviceRequests"`
	Devices              interface{}   `json:"Devices"`
	DNS                  interface{}   `json:"Dns"`
	DNSOptions           interface{}   `json:"DnsOptions"`
	DNSSearch            interface{}   `json:"DnsSearch"`
	Env                  []string      `json:"Env"`
	ExtraHosts           []interface{} `json:"ExtraHosts"`
	GroupAdd             interface{}   `json:"GroupAdd"`
	IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
	IOMaximumIOps        int           `json:"IOMaximumIOps"`
	IpcMode              string        `json:"IpcMode"`
	Isolation            string        `json:"Isolation"`
	KernelMemory         int           `json:"KernelMemory"`
	KernelMemoryTCP      int           `json:"KernelMemoryTCP"`
	Links                interface{}   `json:"Links"`
	LogConfig            interface{}   `json:"LogConfig"`
	MaskedPaths          []string      `json:"MaskedPaths"`
	Memory               int           `json:"Memory"`
	MemoryReservation    int           `json:"MemoryReservation"`
	MemorySwap           int           `json:"MemorySwap"`
	MemorySwappiness     interface{}   `json:"MemorySwappiness"`
	Mounts               []struct {
		ReadOnly bool   `json:"ReadOnly"`
		Source   string `json:"Source"`
		Target   string `json:"Target"`
		Type     string `json:"Type"`
	} `json:"Mounts"`
	NanoCpus       int         `json:"NanoCpus"`
	NetworkMode    string      `json:"NetworkMode"`
	OomKillDisable bool        `json:"OomKillDisable"`
	OomScoreAdj    int         `json:"OomScoreAdj"`
	PidMode        string      `json:"PidMode"`
	PidsLimit      interface{} `json:"PidsLimit"`
	PortBindings   map[string][]struct {
		HostIP   string `json:"HostIp"`
		HostPort string `json:"HostPort"`
	} `json:"PortBindings"`
	Privileged      bool     `json:"Privileged"`
	PublishAllPorts bool     `json:"PublishAllPorts"`
	ReadonlyPaths   []string `json:"ReadonlyPaths"`
	ReadonlyRootfs  bool     `json:"ReadonlyRootfs"`
	RestartPolicy   struct {
		MaximumRetryCount int    `json:"MaximumRetryCount"`
		Name              string `json:"Name"`
	} `json:"RestartPolicy"`
	Runtime     string `json:"Runtime"`
	SecurityOpt interface {
	} `json:"SecurityOpt"`
	ShmSize int    `json:"ShmSize"`
	UTSMode string `json:"UTSMode"`
	Ulimits interface {
	} `json:"Ulimits"`
	UsernsMode   string `json:"UsernsMode"`
	VolumeDriver string `json:"VolumeDriver"`
	VolumesFrom  interface {
	} `json:"VolumesFrom"`
}

type Details struct {
	AppArmorProfile string          `json:"AppArmorProfile"`
	Args            []interface{}   `json:"Args"`
	Config          Config          `json:"Config"`
	Created         time.Time       `json:"Created"`
	Driver          string          `json:"Driver"`
	ExecIDs         interface{}     `json:"ExecIDs"`
	HostConfig      HostConfig      `json:"HostConfig"`
	HostnamePath    string          `json:"HostnamePath"`
	HostsPath       string          `json:"HostsPath"`
	ID              string          `json:"Id"`
	Image           string          `json:"Image"`
	LogPath         string          `json:"LogPath"`
	MountLabel      string          `json:"MountLabel"`
	Mounts          []Mount         `json:"Mounts"`
	Name            string          `json:"Name"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
	Path            string          `json:"Path"`
	Platform        string          `json:"Platform"`
	ProcessLabel    string          `json:"ProcessLabel"`
	ResolvConfPath  string          `json:"ResolvConfPath"`
	RestartCount    int             `json:"RestartCount"`
	State           State           `json:"State"`
	ExeCmd          string          `json:"exe_cmd"`
	FinishTime      int64           `json:"finish_time"`
	Memory          int             `json:"memory"`
	MemoryPercent   float64         `json:"memoryPercent"`
	Status          string          `json:"status"`
	UpTime          int             `json:"up_time"`
}

type DetailsAndProfile struct {
	Details Details `json:"details"`
	Profile Profile `json:"profile"`
}
