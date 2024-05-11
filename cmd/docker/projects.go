package docker

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker"
	"github.com/spf13/cobra"
)

func projectsCmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "projects",
		Aliases: []string{"container"},
		Long: `
The projects command lists projects
`,
	}

	cmd.AddCommand(projectsListCmd(newClient))
	return cmd
}

func projectsListCmd(newClient func() api.Client) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewProjectClient(apiClient).List()
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
			projects := response.Data
			fmt.Printf("found: %v containers\n\n", len(projects))

			fmt.Fprintf(w, "%s\t%s\n", []any{"Name", "Status"}...)
			for _, project := range response.Data {
				fmt.Fprintf(w, "%v\t%v\n", project.Name, project.Status)
			}
			return w.Flush()
		},
	}
}
