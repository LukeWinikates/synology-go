//go:build integration

package images

import (
	"os"
	"testing"

	"github.com/LukeWinikates/synology-go/pkg/api/auth"
	"github.com/stretchr/testify/assert"
)

func TestImagesLifecycle(t *testing.T) {
	loginClient := auth.NewPasswordLoginClient(os.Getenv("DSM_HOST"))
	_, err := loginClient.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)

	ic := NewClient(loginClient.NewAPIClient())

	list, err := ic.List()
	assert.NoError(t, err)
	assert.NotEmpty(t, list.Data.Images)

	img, err := ic.Get("hello-world", "latest")
	assert.NoError(t, err)
	assert.NotEmpty(t, img.Data.Image)
}
