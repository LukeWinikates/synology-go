package containers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"time"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
)

func listHandler(writer http.ResponseWriter, _ *http.Request) {
	internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*ContainerList]{
		Data: &ContainerList{
			Containers: []Container{
				{
					Labels:              nil,
					NetworkSettings:     nil,
					State:               nil,
					Cmd:                 "/bin/bash",
					Created:             0,
					EnableServicePortal: false,
					Exporting:           false,
					FinishTime:          0,
					ID:                  "1234",
					Image:               "foo:latest",
					IsDdsm:              false,
					IsPackage:           false,
					Name:                "Something",
					Services:            nil,
					Status:              "",
					UpStatus:            "",
					UpTime:              0,
				},
			},
			ListResponse: api.ListResponse{
				Limit:  20,
				Offset: 0,
				Total:  1,
			},
		},
		Error:   nil,
		Success: true,
	}))
}

func startServer() *api.TestServer {
	s := api.FakeServerHandler()
	s.AddRoute("GET", APISynoDockerContainer, "list", listHandler)
	s.AddRoute("GET", APISynoDockerContainer, "get", getHandler)
	s.AddRoute("GET", APISynoDockerContainer, "restart", lifecycleHandler)
	s.AddRoute("GET", APISynoDockerContainer, "start", lifecycleHandler)
	s.AddRoute("GET", APISynoDockerContainer, "stop", lifecycleHandler)
	s.AddRoute("GET", APISynoDockerContainerLog, "get", logsHandler)
	s.HTTPServer = httptest.NewServer(s)
	return s
}

func logsHandler(writer http.ResponseWriter, request *http.Request) {
	logLines := []Log{
		{
			Created: time.Time{},
			DocID:   "",
			Stream:  "stdout",
			Text:    "1234",
		},
	}
	limit, err := strconv.Atoi(request.URL.Query().Get("limit"))
	internal.Must(err)
	internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*ContainerLogsResponse]{
		Data: &ContainerLogsResponse{
			ListResponse: api.ListResponse{
				Limit:  limit,
				Offset: 0,
				Total:  1,
			},
			Logs: logLines,
		},
		Error:   nil,
		Success: true,
	}))
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*DetailsAndProfile]{
		Data: &DetailsAndProfile{
			Details: Details{
				AppArmorProfile: "",
				Args:            nil,
				Config:          Config{},
				Created:         time.Time{},
				Driver:          "",
				ExecIDs:         nil,
				HostConfig:      HostConfig{},
				HostnamePath:    "",
				HostsPath:       "",
				ID:              "",
				Image:           "",
				LogPath:         "",
				MountLabel:      "",
				Mounts:          nil,
				Name:            request.URL.Query().Get("name"),
				NetworkSettings: NetworkSettings{},
				Path:            "",
				Platform:        "",
				ProcessLabel:    "",
				ResolvConfPath:  "",
				RestartCount:    0,
				State:           State{},
				ExeCmd:          "",
				FinishTime:      0,
				Memory:          0,
				MemoryPercent:   0,
				Status:          "",
				UpTime:          0,
			},
			Profile: Profile{},
		},
		Error:   nil,
		Success: true,
	}))
}

func lifecycleHandler(writer http.ResponseWriter, request *http.Request) {
	internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*ContainerStats]{
		Data: &ContainerStats{
			CPU:           0,
			Memory:        0,
			MemoryPercent: 0,
			Name:          request.URL.Query().Get("name"),
		},
		Error:   nil,
		Success: true,
	}))
}
