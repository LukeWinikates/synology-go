package projects

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/docker/projects"
	"github.com/spf13/cobra"
)

func projectsStopCmd(builder commandBuilder) *cobra.Command {
	short := "Stop the project with the provided name or id"
	return idRequiredCommand(builder, "stop", short, func(client projects.Client, id string) error {
		return client.Stop(id, func(s string) {
			fmt.Println(s)
		})
	})
}
func projectsStartCmd(builder commandBuilder) *cobra.Command {
	short := "Start the project with the provided name or id"
	return idRequiredCommand(builder, "start", short, func(client projects.Client, id string) error {
		return client.Start(id, func(s string) {
			fmt.Println(s)
		})
	})
}

func projectsBuildCmd(builder commandBuilder) *cobra.Command {
	short := "Build the project with the provided name or id"
	return idRequiredCommand(builder, "build", short, func(client projects.Client, id string) error {
		return client.BuildStream(id, func(s string) {
			fmt.Println(s)
		})
	})
}
