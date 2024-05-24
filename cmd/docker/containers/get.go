package containers

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/containers"
	"github.com/spf13/cobra"
)

func getCmd(newClient func() api.Client) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := containers.NewClient(apiClient).GetContainer(name)
			if err != nil {
				return err
			}

			fmt.Println(response)
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))
	return cmd
}
