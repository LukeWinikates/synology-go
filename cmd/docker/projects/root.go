package projects

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/projects"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project"},
	}
	factory := func() projects.Client { return projects.NewClient(newClient()) }
	builder := commandBuilder{
		clientFactory: factory,
		objectPrinter: internal.NewTableWriter[*projects.Project]([]string{
			"Name",
			"ID",
			"Status",
		}, func(item *projects.Project) []string {
			return []string{item.Name, item.ID, item.Status}
		}),
	}

	cmd.AddCommand(listCmd(builder))
	cmd.AddCommand(projectsGetCmd(builder))
	cmd.AddCommand(projectsCreateCmd(builder))
	cmd.AddCommand(projectsUpdateCmd(builder))
	cmd.AddCommand(projectsStopCmd(builder))
	cmd.AddCommand(projectsStartCmd(builder))
	cmd.AddCommand(projectsBuildCmd(builder))
	return cmd
}

type commandBuilder struct {
	clientFactory func() projects.Client
	objectPrinter internal.TableWriter[*projects.Project]
}

func idRequiredCommand(builder commandBuilder, use string, commandFunc func(client projects.Client, id string) error) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use: use,
		RunE: func(_ *cobra.Command, _ []string) error {
			return commandFunc(builder.clientFactory(), id)
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "project id")
	internal.Must(cmd.MarkFlagRequired("id"))
	return cmd
}
