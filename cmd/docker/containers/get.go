package containers

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/docker/containers"
	"github.com/spf13/cobra"
)

func getCmd(builder commandBuilder) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().GetContainer(name)
			if err != nil {
				return err
			}
			commandRunner.Client().GetContainer(name)
			commandRunner.WriteOutput(response.Data)
			writer := internal.NewTableWriter[*containers.DetailsAndProfile]([]string{"Name"}, func(item *containers.DetailsAndProfile) []string {
				return []string{item.Profile.Name}
			})
			err = writer.Write(response.Data)
			if err != nil {
				return err
			}
			return writer.Flush()
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))
	return cmd
}
