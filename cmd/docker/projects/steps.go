package projects

import (
	"fmt"

	"github.com/LukeWinikates/synology-go/pkg/docker/projects"
)

func buildProject(client projects.Client, projectID string) error {
	fmt.Println("following build output...")
	err := client.BuildStream(projectID, func(s string) {
		fmt.Println(s)
	})
	if err != nil {
		return err
	}
	fmt.Println("build completed")
	return nil
}

func stopProject(client projects.Client, projectID string) error {
	fmt.Println("stopping project...")
	err := client.Stop(projectID, func(s string) {
		fmt.Println(s)
	})
	if err != nil {
		return fmt.Errorf("couldn't stop project (%s)", err.Error())
	}
	return err
}

func findProjectByName(client projects.Client, name string) (*projects.Project, error) {
	list, err := client.List()
	if err != nil {
		return nil, err
	}
	for _, project := range list.Data {
		if project.Name == name {
			return &project, nil
		}
	}
	return nil, fmt.Errorf("project not found: %s", name)
}
