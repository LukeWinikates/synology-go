package containers

import (
	"github.com/spf13/cobra"
)

func listCmd(builder commandBuilder) *cobra.Command {
	return &cobra.Command{
		Use:  "list",
		Long: "",
		RunE: func(_ *cobra.Command, _ []string) error {
			response, err := commandRunner.Client().ListContainers()
			if err != nil {
				return err
			}

			commandRunner.Progress.Printf("\"found: %v containers\\n\\n\", response.Data.Total")
			return commandRunner.Output.Write(response.Data.Containers)
			//fmt.Printf("found: %v containers\n\n", response.Data.Total)
			//for _, container := range response.Data.Containers {
			//	err = builder.containerWriter.Write(container)
			//	if err != nil {
			//		return err
			//	}
			//}
			//return builder.containerWriter.Flush()
		},
	}
}
