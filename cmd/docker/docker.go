package docker

import (
	"fmt"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	containerCmd := &cobra.Command{
		Use: "container",
	}

	containerCmd.AddCommand(listCmd(newClient))
	containerCmd.AddCommand(restartCmd(newClient))
	containerCmd.AddCommand(getCmd(newClient))

	dockerCmd := &cobra.Command{
		Use: "docker",
	}

	dockerCmd.AddCommand(containerCmd)

	return dockerCmd
}

func getCmd(newClient func() api.Client) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).GetContainer(name)
			if err != nil {
				return err
			}

			fmt.Println(response)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&name, "name", "", "container name")

	return cmd
}

func restartCmd(newClient func() api.Client) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:  "restart",
		Long: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).RestartContainer(name)
			if err != nil {
				return err
			}

			if response.Success {
				fmt.Printf("Container \"%s\" restarted\n", name)
			} else {
				return fmt.Errorf("could not restart container: %s", name)
			}

			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&name, "name", "", "container name")

	return cmd
}

func listCmd(newClient func() api.Client) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).ListContainers()
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
			fmt.Printf("found: %v containers\n\n", response.Data.Total)

			fmt.Fprintf(w, "%s\t%s\n", []any{"Name", "Status"}...)
			for _, container := range response.Data.Containers {
				fmt.Fprintf(w, "%v\t%v\n", container.Name, container.Status)
			}
			return w.Flush()
		},
	}
}
