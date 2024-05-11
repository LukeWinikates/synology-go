//go:build integration
// +build integration

package docker

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const COMPOSE_FILE_CONTENTS = `version: '3'
services:
  hello:
    image: hello-world
    container_name: hello-world-test1`

func TestProjectLifecycle(t *testing.T) {
	c, _ := api.NewClient(
		os.Getenv("DSM_HOST"))
	_, err := c.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)

	pc := NewProjectClient(c)
	projectListResponse, err := pc.List()
	assert.NoError(t, err)

	for _, v := range projectListResponse.Data {
		if v.Name == "test1" {
			assert.Fail(t, "found existing project using test fixture name", v.Name, v.ID)
		}
	}

	_, err = pc.GetShareInfo("/test/test1")
	assert.NoError(t, err)

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
