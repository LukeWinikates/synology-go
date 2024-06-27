//go:build integration
// +build integration

package projects

import (
	"fmt"
	"github.com/LukeWinikates/synology-go/pkg/api/auth"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const COMPOSE_FILE_CONTENTS = `
version: '3'
services:
  hello:
    image: hello-world
    container_name: hello-world-test1`
const COMPOSE_FILE_CONTENTS_UPDATE = `
version: '3'
services:
  hello:
    image: hello-world
    container_name: hello-world-test1
  hello2:
    image: hello-world
    container_name: "hello-world-test2"
  hello3:
    image: hello-world
    container_name: hello-world-test3
  hello4:
    image: hello-world
    container_name: hello-world-test4`

func TestProjectLifecycle(t *testing.T) {
	loginClient := auth.NewPasswordLoginClient(os.Getenv("DSM_HOST"))
	_, err := loginClient.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)
	pc := NewClient(loginClient.NewAPIClient())

	assert.NoError(t, checkPreconditions(pc, "test1", "/test/test1"))

	createResponse, err := pc.Create("test1", "/test/test1", COMPOSE_FILE_CONTENTS, nil)
	assert.NoError(t, err)
	projectId := createResponse.Data.ID

	buildStreamText := ""
	err = pc.BuildStream(projectId, func(s string) {
		buildStreamText += s
	})
	assert.NoError(t, err)
	assert.Contains(t, buildStreamText, "Exit Code: 0")

	deleteResponse, err := pc.Delete(projectId)

	assert.NoError(t, err)
	assert.True(t, deleteResponse.Success)
}

func TestProjectUpdateComposeYaml(t *testing.T) {
	loginClient := auth.NewPasswordLoginClient(os.Getenv("DSM_HOST"))
	_, err := loginClient.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)
	pc := NewClient(loginClient.NewAPIClient())

	assert.NoError(t, checkPreconditions(pc, "test1", "/test/test1"))

	createResponse, err := pc.Create("test1", "/test/test1", COMPOSE_FILE_CONTENTS, nil)
	assert.NoError(t, err)
	projectId := createResponse.Data.ID

	buildStreamText := ""
	err = pc.BuildStream(projectId, func(s string) {
		buildStreamText += s
	})
	assert.NoError(t, err)
	assert.Contains(t, buildStreamText, "Exit Code: 0")

	updateResponse, err := pc.UpdateContent(projectId, COMPOSE_FILE_CONTENTS_UPDATE)
	assert.NoError(t, err)
	assert.True(t, updateResponse.Success)

	buildStreamText = ""
	err = pc.BuildStream(projectId, func(s string) {
		buildStreamText += s
	})
	assert.NoError(t, err)
	assert.Contains(t, buildStreamText, "Exit Code: 0")
	deleteResponse, err := pc.Delete(projectId)

	assert.NoError(t, err)
	assert.True(t, deleteResponse.Success)
}

func checkPreconditions(pc Client, projectName string, sharePath string) error {
	projectListResponse, err := pc.List()
	if err != nil {
		return err
	}

	for _, v := range projectListResponse.Data {
		if v.Name == projectName {
			return fmt.Errorf("found existing project using test fixture name: (%s %s)", v.Name, v.ID)
		}
	}

	_, err = pc.GetShareInfo(sharePath)
	return err
}
