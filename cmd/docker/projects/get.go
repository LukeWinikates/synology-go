package projects

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/docker/projects"
	"github.com/spf13/cobra"
)

func projectsGetCmd(builder commandBuilder) *cobra.Command {
	return idRequiredCommand(builder, "get", func(client projects.Client, id string) error {
		response, err := client.Get(id)
		if err != nil {
			return err
		}

		err = builder.objectPrinter.Write(response.Data)
		if err != nil {
			return err
		}
		return builder.objectPrinter.Flush()
	})
}

func projectsDeleteCmd(builder commandBuilder) *cobra.Command {
	return idRequiredCommand(builder, "delete", func(client projects.Client, id string) error {
		response, err := client.Delete(id)
		if err != nil {
			return err
		}
		fmt.Println(response.Data)
		return err
	})
}
