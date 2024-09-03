package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateRepositoryName(t *testing.T) {
	require.NoError(t, ValidateRepositoryName("foo/bar"))
	expectedError := ValidateRepositoryName("repository.foo.local/foo/bar")
	require.Error(t, expectedError)
	assert.Contains(t, expectedError.Error(), "maybe you want 'foo/bar'?")
}

func TestRepositoryShortName(t *testing.T) {
	assert.Equal(t, "foo/bar", RepositoryShortName("foo/bar"))
	assert.Equal(t, "foo/bar", RepositoryShortName("repository.foo.local/foo/bar"))
}
