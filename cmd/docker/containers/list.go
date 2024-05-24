package containers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/containers"
	"github.com/spf13/cobra"
)

func listCmd(newClient func() api.Client) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := containers.NewClient(apiClient).ListContainers()
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
			fmt.Printf("found: %v containers\n\n", response.Data.Total)

			fmt.Fprintf(w, "%s\t%s\n", []any{"Name", "Status"}...)
			for _, container := range response.Data.Containers {
				fmt.Fprintf(w, "%v\t%v\n", container.Name, container.Status)
			}
			return w.Flush()
		},
	}
}
