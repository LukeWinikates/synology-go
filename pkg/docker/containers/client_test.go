package containers

import (
	"testing"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListContainers(t *testing.T) {
	server, _ := startServer()
	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.ListContainers()
	require.NoError(t, err)
	assert.NotEmpty(t, containers.Data.Containers)
	defer server.Close()
}

func TestGetContainer(t *testing.T) {
	server, _ := startServer()

	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.GetContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Details.Name)
	defer server.Close()
}

func TestStartContainer(t *testing.T) {
	server, _ := startServer()
	defer server.Close()
	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.StartContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}

func TestStopContainer(t *testing.T) {
	server, _ := startServer()
	defer server.Close()
	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.StopContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}

func TestRestartContainer(t *testing.T) {
	server, _ := startServer()
	defer server.Close()

	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.RestartContainer("hello-world")
	require.NoError(t, err)
	assert.Equal(t, "hello-world", containers.Data.Name)
}
