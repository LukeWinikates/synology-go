package images

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/docker/images"
	"github.com/spf13/cobra"
)

func listCmd(clientFactory func() images.Client) *cobra.Command {
	return &cobra.Command{
		Use: "list",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := clientFactory().List()
			if err != nil {
				return err
			}
			fmt.Printf("found: %v images\n", len(response.Data.Images))
			printer := listImagePrinter()
			for _, image := range response.Data.Images {
				err = printer.Write(&image)
				if err != nil {
					return err
				}
			}
			return printer.Flush()
		},
	}
}
