package containers

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func containerLogsCmd(builder commandBuilder) *cobra.Command {
	var name string
	var simple bool
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Print the logs of the container with the provided name",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().GetContainerLogs(name)
			if err != nil {
				return err
			}

			for _, log := range response.Data.Logs {
				if simple {
					fmt.Printf("%s", log.Text)
				} else {
					fmt.Printf("[%s] %s: %s", log.Stream, log.Created, log.Text)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))
	cmd.Flags().BoolVarP(&simple, "simple", "s", false, "only print log content")

	return cmd
}
