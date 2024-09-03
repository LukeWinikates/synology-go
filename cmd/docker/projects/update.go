package projects

import (
	"fmt"
	"os"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func projectsUpdateCmd(builder commandBuilder) *cobra.Command {
	var name string
	var filePath string
	var autoBuild bool
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Modify and rebuild the compose file for the projects with the given name",
		RunE: func(_ *cobra.Command, _ []string) error {
			client := builder.clientFactory()
			content, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}
			project, err := findProjectByName(client, name)
			if err != nil {
				return fmt.Errorf("couldn't find project with name (%s)", err.Error())
			}
			err = stopProject(client, project.ID)
			if err != nil {
				return err
			}

			response, err := client.UpdateContent(project.ID, string(content))
			if err != nil {
				return fmt.Errorf("couldn't update project (%s)", err.Error())
			}

			err = builder.objectPrinter.Write(response.Data)
			if err != nil {
				return err
			}
			err = builder.objectPrinter.Flush()
			if err != nil {
				return err
			}
			if autoBuild {
				return buildProject(client, project.ID)
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "project name")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "compose file path")
	internal.Must(cmd.MarkFlagRequired("name"))
	internal.Must(cmd.MarkFlagRequired("file"))
	cmd.Flags().BoolVarP(&autoBuild, "build", "b", false, "automatically build project")

	return cmd
}
