package images

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/docker/images"
	"github.com/spf13/cobra"
)

func getCmd(clientFactory func() images.Client) *cobra.Command {
	var name, tag string
	cmd := &cobra.Command{
		Use: "get",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := clientFactory().Get(name, tag)
			if err != nil {
				return err
			}

			printer := imagePrinter()
			err = printer.Write(response.Data)
			if err != nil {
				return err
			}
			return printer.Flush()
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "image name")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "image tag")
	internal.Must(cmd.MarkFlagRequired("name"))
	internal.Must(cmd.MarkFlagRequired("tag"))
	return cmd
}
