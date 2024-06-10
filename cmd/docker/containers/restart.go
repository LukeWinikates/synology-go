package containers

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func restartCmd(builder commandBuilder) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:  "restart",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().RestartContainer(name)
			if err != nil {
				return err
			}

			if response.Success {
				fmt.Printf("Container \"%s\" restarted\n", name)
			} else {
				return fmt.Errorf("could not restart container: %s", name)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))

	return cmd
}
