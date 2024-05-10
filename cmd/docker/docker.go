package docker

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	containerCmd := &cobra.Command{
		Use:     "container",
		Aliases: []string{"containers"},
		Long: `
The containers command lists, describes, restarts, and shows logs for containers
`,
	}

	containerCmd.AddCommand(listCmd(newClient))
	containerCmd.AddCommand(restartCmd(newClient))
	containerCmd.AddCommand(getCmd(newClient))
	containerCmd.AddCommand(containerLogsCmd(newClient))

	dockerCmd := &cobra.Command{
		Use: "docker",
		Long: `
The docker command communicates with the Container Manager application. 
Use it to examine, restart, or delete containers, or to view logs from your containers.`,
	}

	dockerCmd.AddCommand(containerCmd)
	dockerCmd.AddCommand(logsCmd(newClient))

	return dockerCmd
}

func getCmd(newClient func() api.Client) *cobra.Command {
	var name string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).GetContainer(name)
			if err != nil {
				return err
			}

			fmt.Println(response)
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "", "container name")
	must(cmd.MarkFlagRequired("name"))
	return cmd
}

func must(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func containerLogsCmd(newClient func() api.Client) *cobra.Command {
	var name string
	var simple bool
	cmd := &cobra.Command{
		Use:  "logs",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).GetContainerLogs(name)
			if err != nil {
				return err
			}

			for _, log := range response.Data.Logs {
				if simple {
					fmt.Printf("%s", log.Text)
				} else {
					fmt.Printf("[%s] %s: %s", log.Stream, log.Created, log.Text)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "container name")
	must(cmd.MarkFlagRequired("name"))
	cmd.Flags().BoolVarP(&simple, "simple", "s", false, "only print log content")

	return cmd
}
func logsCmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use: "logs",
		Long: `
The logs command lists the recent container lifecycle events logged by the Container Manger application

If you are looking for container logs, see: dsmctl docker container logs --name $CONTAINER_NAME
`,
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewClient(apiClient).GetContainerManagerLogs()
			if err != nil {
				return err
			}

			for _, logLine := range response.Data.Logs {
				fmt.Printf("[%s] %s: %s\n", logLine.Level, logLine.Time, logLine.Event)
			}
			return nil
		},
	}

	return cmd
}

func restartCmd(newClient func() api.Client) *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:  "restart",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
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
	cmd.Flags().StringVar(&name, "name", "", "container name")
	must(cmd.MarkFlagRequired("name"))

	return cmd
}

func listCmd(newClient func() api.Client) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
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
