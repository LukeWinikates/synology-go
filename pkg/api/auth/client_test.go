package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newServerHandler(account, password string) *api.TestServer {
	handler := api.FakeServerHandler()
	handler.AddRoute("GET", "SYNO.API.Auth", "login", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		if query["passwd"][0] == password && query["account"][0] == account {
			internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*LoginResponse]{
				Data: &LoginResponse{
					Did:          "20102",
					IsPortalPort: false,
					Sid:          "414141",
				},
				Error:   nil,
				Success: true,
			}))
		} else {
			internal.Must(json.NewEncoder(writer).Encode(&api.ResponseWrapper[*LoginResponse]{
				Data:    nil,
				Error:   map[string]interface{}{"code": 400},
				Success: false,
			}))
		}
	})
	return handler

}

func TestLogin(t *testing.T) {
	s := httptest.NewServer(newServerHandler("some login", "some password"))
	defer s.Close()
	login := NewPasswordLoginClient(s.URL)
	session, err := login.Login(
		"some login",
		"some password")
	require.NoError(t, err)
	assert.True(t, session.Success)
	assert.NotEmpty(t, session.Data.Sid)
}

func TestFailedLogin(t *testing.T) {
	s := httptest.NewServer(newServerHandler("some login", "some password"))
	defer s.Close()
	login := NewPasswordLoginClient(
		s.URL)
	session, err := login.Login(
		"some login",
		"some wrong password")
	require.Error(t, err)
	assert.Nil(t, session)
	assert.Contains(t, err.Error(), "failed request: 400: No such account or incorrect password")
}
