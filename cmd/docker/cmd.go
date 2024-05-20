package docker

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	dockerCmd := &cobra.Command{
		Use: "docker",
		Long: `
The docker command communicates with the Container Manager application. 
Use it to examine, restart, or delete containers, or to view logs from your containers.`,
	}
	dockerCmd.AddCommand(logsCmd(newClient))

	dockerCmd.AddCommand(containerCmd(newClient))
	dockerCmd.AddCommand(projectsCmd(newClient))

	return dockerCmd
}

func logsCmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use: "logs",
		Long: `
The logs command lists the recent container lifecycle events logged by the Container Manger application

If you are looking for container logs, see: dsmctl docker container logs --name $CONTAINER_NAME
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).GetContainerManagerLogs()
			if err != nil {
				return err
			}

			for _, logLine := range response.Data.Logs {
				fmt.Printf("[%s] %s: %s\n", logLine.Level, logLine.Time, logLine.Event)
			}
			return nil
		},
	}

	return cmd
}
