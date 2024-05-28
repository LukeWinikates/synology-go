//go:build integration
// +build integration

package images

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImagesLifecycle(t *testing.T) {
	c, _ := api.NewClient(
		os.Getenv("DSM_HOST"))
	_, err := c.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)

	ic := NewClient(c)

	list, err := ic.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, list.Data.Images)

	img, err := ic.Get("hello-world", "latest")
	assert.NoError(t, err)
	assert.NotEmpty(t, img.Data.Image)
}
