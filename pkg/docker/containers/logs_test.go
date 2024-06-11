package containers

import (
	"testing"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogs(t *testing.T) {
	server, testServer := startServer()
	defer server.Close()
	newClient, err := api.NewClient(server.URL)
	require.NoError(t, err)
	c := NewClient(newClient)
	containers, err := c.GetContainerLogs("hello-world")
	require.NoError(t, err)
	assert.Len(t, containers.Data.Logs, 1)
	assert.Equal(t, 1, containers.Data.Total)
	assert.Equal(t, 100, containers.Data.Limit)
	assert.Equal(t, `"hello-world"`, testServer.Requests[0].URL.Query().Get("name"))
}
