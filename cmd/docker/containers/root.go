package containers

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/containers"
	"github.com/spf13/cobra"
)

type commandBuilder struct {
	newClient       func() containers.Client
	containerWriter internal.TableWriter[containers.Container]
}

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "container",
		Aliases: []string{"containers"},
		Short:   `List, describe, restart, and shows logs for containers`,
	}

	builder := commandBuilder{
		newClient: func() containers.Client {
			return containers.NewClient(newClient())
		},
		containerWriter: internal.NewTableWriter[containers.Container](
			[]string{"Name", "Status"}, func(item containers.Container) []string {
				return []string{
					item.Name,
					item.Status,
				}
			}),
	}

	cmd.AddCommand(listCmd(builder))
	cmd.AddCommand(restartCmd(builder))
	cmd.AddCommand(startCmd(builder))
	cmd.AddCommand(stopCmd(builder))
	cmd.AddCommand(getCmd(builder))
	cmd.AddCommand(containerLogsCmd(builder))
	return cmd
}
