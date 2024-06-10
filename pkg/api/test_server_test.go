package api

import (
	"encoding/json"
	"net/http"

	"github.com/LukeWinikates/synology-go/internal"
)

type testServer struct {
	password string
	account  string
}

func (t testServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/webapi/entry.cgi" {
		writer.WriteHeader(http.StatusNotFound)
	}

	if request.Method == "GET" {
		query := request.URL.Query()
		if query.Get("method") == "login" && query.Get("api") == "SYNO.API.Auth" {
			if query["passwd"][0] == t.password && query["account"][0] == t.account {
				internal.Must(json.NewEncoder(writer).Encode(&ResponseWrapper[*LoginResponse]{
					Data: &LoginResponse{
						Did:          "20102",
						IsPortalPort: false,
						Sid:          "414141",
					},
					Error:   nil,
					Success: true,
				}))
			} else {
				internal.Must(json.NewEncoder(writer).Encode(&ResponseWrapper[*LoginResponse]{
					Data:    nil,
					Error:   map[string]interface{}{"code": 400},
					Success: false,
				}))
			}
		}

	}

}

func FakeServerHandler(account string, password string) http.Handler {
	return &testServer{account: account, password: password}
}
