//go:build integration
// +build integration

package auth

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoginIntegration(t *testing.T) {
	c := NewPasswordLoginClient(
		os.Getenv("DSM_HOST"))
	session, err := c.Login(
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	assert.NoError(t, err)
	assert.True(t, session.Success)
	assert.NotEmpty(t, session.Data.Sid)
}
