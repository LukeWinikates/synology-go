package api

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLogin(t *testing.T) {
	s := httptest.NewServer(FakeServerHandler("some login", "some password"))
	defer s.Close()
	c, _ := NewClient(
		s.URL)
	session, err := c.Login(
		"some login",
		"some password")
	require.NoError(t, err)
	assert.True(t, session.Success)
	assert.NotEmpty(t, session.Data.Sid)
}

func TestFailedLogin(t *testing.T) {
	s := httptest.NewServer(FakeServerHandler("some login", "some password"))
	defer s.Close()
	c, _ := NewClient(
		s.URL)
	session, err := c.Login(
		"some login",
		"some wrong password")
	require.Error(t, err)
	assert.Nil(t, session)
	assert.Contains(t, err.Error(), "failed request: 400: No such account or incorrect password")
}
