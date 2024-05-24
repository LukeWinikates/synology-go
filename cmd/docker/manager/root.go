package manager

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/manager"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "manager",
		Long: `The manager shows logs for the container manager app`,
	}

	cmd.AddCommand(logsCmd(newClient))
	return cmd
}

func logsCmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use: "logs",
		Long: `
The logs command lists the recent container lifecycle events logged by the Container Manger application

If you are looking for container logs, see: synoctl docker container logs --name $CONTAINER_NAME
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := manager.NewClient(newClient()).GetContainerManagerLogs()
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
