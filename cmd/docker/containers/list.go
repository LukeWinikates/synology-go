package containers

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listCmd(builder commandBuilder) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all containers",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().ListContainers()
			if err != nil {
				return err
			}
			fmt.Printf("found: %v containers\n\n", response.Data.Total)
			for _, container := range response.Data.Containers {
				err = builder.containerWriter.Write(container)
				if err != nil {
					return err
				}
			}
			return builder.containerWriter.Flush()
		},
	}
}
