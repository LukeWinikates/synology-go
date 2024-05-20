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
	cmd.AddCommand(projectsGetCmd(newClient))
	cmd.AddCommand(projectsCreateCmd(newClient))
	//cmd.AddCommand(projectsUpdateCmd(newClient))
	return cmd
}

func projectsGetCmd(newClient func() api.Client) *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:  "get",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			response, err := docker.NewProjectClient(apiClient).Get(id)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
			fmt.Fprintf(w, "%s\t%s\n", []any{"Name", "Status"}...)
			fmt.Fprintf(w, "%v\t%v\n", response.Data.Name, response.Data.Status)
			return w.Flush()
		},
	}
	cmd.Flags().StringVarP(&id, "id", "i", "", "project id")
	must(cmd.MarkFlagRequired("id"))

	return cmd
}
func projectsCreateCmd(newClient func() api.Client) *cobra.Command {
	var name string
	var sharePath string
	var filePath string
	var autoBuild bool
	cmd := &cobra.Command{
		Use:  "create",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			apiClient := newClient()
			content, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}
			client := docker.NewProjectClient(apiClient)
			response, err := client.Create(name, sharePath, string(content), nil)
			if err != nil {
				return err
			}

			w := tabwriter.NewWriter(os.Stdout, 1, 8, 1, ' ', 0)
			fmt.Fprintf(w, "%s\t%s\n", []any{"Name", "Status"}...)
			fmt.Fprintf(w, "%v\t%v\n", response.Data.Name, response.Data.Status)
			w.Flush()
			if autoBuild {
				fmt.Println("following build output...")
				return client.BuildStream(response.Data.ID, func(s string) {
					fmt.Println(s)
				})
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "i", "", "project name")
	cmd.Flags().StringVarP(&sharePath, "share-path", "s", "", "path on Synology host")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "compose file path")
	must(cmd.MarkFlagRequired("name"))
	must(cmd.MarkFlagRequired("share-path"))
	must(cmd.MarkFlagRequired("file"))
	cmd.Flags().BoolVarP(&autoBuild, "build", "b", false, "automatically build project")

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
