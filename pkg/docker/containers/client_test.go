package containers

import (
	"testing"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/api/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListContainers(t *testing.T) {
	testServer := startServer()
	newClient := api.NewClient(testServer.HTTPServer.URL, auth.NewSessionAuthorizer(testServer.Session()))
	c := NewClient(newClient)

	containers, err := c.ListContainers()
	require.NoError(t, err)
	assert.NotEmpty(t, containers.Data.Containers)
	defer testServer.HTTPServer.Close()
}

func TestGetContainer(t *testing.T) {
	testServer := startServer()
	newClient := api.NewClient(testServer.HTTPServer.URL, auth.NewSessionAuthorizer(testServer.Session()))

	c := NewClient(newClient)
	containers, err := c.GetContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Details.Name)
	defer testServer.HTTPServer.Close()
}

func TestStartContainer(t *testing.T) {
	testServer := startServer()

	defer testServer.HTTPServer.Close()
	newClient := api.NewClient(testServer.HTTPServer.URL, auth.NewSessionAuthorizer(testServer.Session()))
	c := NewClient(newClient)
	containers, err := c.StartContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}

func TestStopContainer(t *testing.T) {
	testServer := startServer()
	defer testServer.HTTPServer.Close()

	newClient := api.NewClient(testServer.HTTPServer.URL, auth.NewSessionAuthorizer(testServer.Session()))
	c := NewClient(newClient)
	containers, err := c.StopContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}

func TestRestartContainer(t *testing.T) {
	testServer := startServer()
	defer testServer.HTTPServer.Close()

	newClient := api.NewClient(testServer.HTTPServer.URL, auth.NewSessionAuthorizer(testServer.Session()))
	c := NewClient(newClient)
	containers, err := c.RestartContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}
