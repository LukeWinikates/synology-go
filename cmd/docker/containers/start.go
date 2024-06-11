package containers

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func startCmd(builder commandBuilder) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:  "start",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().StartContainer(name)
			if err != nil {
				return err
			}
			if response.Success {
				fmt.Printf("Container \"%s\" started\n", name)
			} else {
				return fmt.Errorf("could not start container: %s", name)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))

	return cmd
}
