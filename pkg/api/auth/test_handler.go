package auth

import (
	"encoding/json"
	"net/http"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
)

func NewFakeAuthServerHandler(account, password string) *api.TestServer {
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
