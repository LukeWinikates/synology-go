package docker

import (
	"github.com/LukeWinikates/synology-go/cmd/docker/containers"
	"github.com/LukeWinikates/synology-go/cmd/docker/images"
	"github.com/LukeWinikates/synology-go/cmd/docker/manager"
	"github.com/LukeWinikates/synology-go/cmd/docker/projects"
	"github.com/LukeWinikates/synology-go/cmd/docker/registries"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage containers, images, projects and registries",
		Long: `
The docker command communicates with the Container Manager application. 
Use it to examine, restart, or delete containers and projects, or to view logs from your containers.`,
	}
	dockerCmd.AddCommand(manager.Cmd(newClient))
	dockerCmd.AddCommand(containers.Cmd(newClient))
	dockerCmd.AddCommand(projects.Cmd(newClient))
	dockerCmd.AddCommand(images.Cmd(newClient))
	dockerCmd.AddCommand(registries.Cmd(newClient))

	return dockerCmd
}
