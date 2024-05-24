package containers

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "container",
		Aliases: []string{"containers"},
		Long:    `The containers command lists, describes, restarts, and shows logs for containers`,
	}

	cmd.AddCommand(listCmd(newClient))
	cmd.AddCommand(restartCmd(newClient))
	cmd.AddCommand(getCmd(newClient))
	cmd.AddCommand(containerLogsCmd(newClient))
	return cmd
}
