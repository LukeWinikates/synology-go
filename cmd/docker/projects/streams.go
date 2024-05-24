package projects

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/docker/projects"
	"github.com/spf13/cobra"
)

func projectsStopCmd(builder commandBuilder) *cobra.Command {
	return idRequiredCommand(builder, "stop", func(client projects.Client, id string) error {
		return client.Stop(id, func(s string) {
			fmt.Println(s)
		})
	})
}
func projectsStartCmd(builder commandBuilder) *cobra.Command {
	return idRequiredCommand(builder, "start", func(client projects.Client, id string) error {
		return client.Start(id, func(s string) {
			fmt.Println(s)
		})
	})
}

func projectsBuildCmd(builder commandBuilder) *cobra.Command {
	return idRequiredCommand(builder, "build", func(client projects.Client, id string) error {
		return client.BuildStream(id, func(s string) {
			fmt.Println(s)
		})
	})
}
