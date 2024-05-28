package images

import "github.com/LukeWinikates/synology-go/pkg/api"

type Port struct {
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
}
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Image struct {
	Author        string        `json:"author"`
	Cmd           []interface{} `json:"cmd"`
	Digest        string        `json:"digest"`
	DockerVersion string        `json:"docker_version"`
	Entrypoint    []string      `json:"entrypoint"`
	Env           []KeyValue    `json:"env"`
	ID            string        `json:"id"`
	Image         string        `json:"image"`
	Ports         []Port        `json:"ports"`
	Size          int           `json:"size"`
	Tag           string        `json:"tag"`
	VirtualSize   int           `json:"virtual_size"`
	Volumes       []interface{} `json:"volumes"`
}

type Task struct {
	TaskID string `json:"task_id"`
}

type UpgradeStatus struct {
	Current  int    `json:"current"`
	Finished bool   `json:"finished"`
	Image    string `json:"image"`
	State    string `json:"state"`
	Total    int    `json:"total"`
}

type ListImage struct {
	Created      int      `json:"created"`
	Description  string   `json:"description"`
	Digest       string   `json:"digest"`
	ID           string   `json:"id"`
	RemoteDigest string   `json:"remote_digest"`
	Repository   string   `json:"repository"`
	Size         int      `json:"size"`
	Tags         []string `json:"tags"`
	Upgradable   bool     `json:"upgradable"`
	VirtualSize  int      `json:"virtual_size"`
}

type ImageList struct {
	api.ListResponse
	Images []ListImage `json:"images"`
}

type PullStatus struct {
	Description string `json:"description"`
	Downloaded  int    `json:"downloaded"`
	Finished    bool   `json:"finished"`
	Repository  string `json:"repository"`
	Tag         string `json:"tag"`
}
