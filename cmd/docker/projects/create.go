package projects

import (
	"fmt"
	"os"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/spf13/cobra"
)

func projectsCreateCmd(builder commandBuilder) *cobra.Command {
	var name string
	var sharePath string
	var filePath string
	var autoBuild bool
	cmd := &cobra.Command{
		Use: "create",
		RunE: func(_ *cobra.Command, _ []string) error {
			content, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}
			client := builder.clientFactory()
			response, err := client.Create(name, sharePath, string(content), nil)
			if err != nil {
				return err
			}
			project := response.Data
			err = builder.objectPrinter.Write(project)
			if err != nil {
				return err
			}
			err = builder.objectPrinter.Flush()
			if err != nil {
				return err
			}
			if autoBuild {
				fmt.Println("following build output...")
				return client.BuildStream(project.ID, func(s string) {
					fmt.Println(s)
				})
			}

			return nil
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "project name")
	cmd.Flags().StringVarP(&sharePath, "share-path", "s", "", "path on Synology host")
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "compose file path")
	internal.Must(cmd.MarkFlagRequired("name"))
	internal.Must(cmd.MarkFlagRequired("share-path"))
	internal.Must(cmd.MarkFlagRequired("file"))
	cmd.Flags().BoolVarP(&autoBuild, "build", "b", false, "automatically build project")

	return cmd
}
