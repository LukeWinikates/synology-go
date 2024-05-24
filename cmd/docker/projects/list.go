package projects

import (
	"fmt"

	"github.com/spf13/cobra"
)

func listCmd(builder commandBuilder) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.clientFactory().List()
			if err != nil {
				return err
			}
			fmt.Printf("found: %v projects\n\n", len(response.Data))
			for _, project := range response.Data {
				err = builder.objectPrinter.Write(&project)
				if err != nil {
					return err
				}
			}
			return builder.objectPrinter.Flush()
		},
	}
}
