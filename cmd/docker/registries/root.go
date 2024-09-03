package registries

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/registries"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use: "registries",
	}

	cmd.AddCommand(listCMD(newClient))
	return cmd
}

func listCMD(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(_ *cobra.Command, _ []string) error {

			response, err := registries.NewClient(newClient()).List()
			if err != nil {
				return err
			}
			writer := internal.NewTableWriter[registries.Registry](
				[]string{"Using", "Name", "URL"},
				func(item registries.Registry) []string {
					isUsing := ""
					if response.Data.Using == item.Name {
						isUsing = "*"
					}
					return []string{isUsing, item.Name, item.URL}
				})
			for _, registry := range response.Data.Registries {
				err = writer.Write(registry)
				if err != nil {
					return err
				}
			}
			return writer.Flush()
		},
	}

	return cmd
}
