package containers

import (
	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func getCmd(builder commandBuilder) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := builder.newClient().GetContainer(name)
			if err != nil {
				return err
			}

			err = builder.containerWriter.Write(*response.Data)
			if err != nil {
				return err
			}
			return builder.containerWriter.Flush()
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	internal.Must(cmd.MarkFlagRequired("name"))
	return cmd
}
