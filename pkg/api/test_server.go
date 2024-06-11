package api

import (
	"encoding/json"
	"net/http"

	"github.com/LukeWinikates/synology-go/internal"
)

type mockRoute struct {
	api        string
	apiMethod  string
	httpMethod string
	handler    http.HandlerFunc
}

type TestServer struct {
	password string
	account  string
	routes   []mockRoute
	Requests []*http.Request
}

func (t *TestServer) AddRoute(httpMethod, api, apiMethod string, handler http.HandlerFunc) {
	t.routes = append(t.routes, mockRoute{
		api:        api,
		apiMethod:  apiMethod,
		httpMethod: httpMethod,
		handler:    handler,
	})
}

func (t *TestServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	t.Requests = append(t.Requests, request)
	if request.URL.Path != "/webapi/entry.cgi" {
		writer.WriteHeader(http.StatusNotFound)
	}

	params := request.URL.Query()
	for _, route := range t.routes {
		if request.Method == http.MethodPost {
			internal.Must(request.ParseForm())
			params = request.Form
		}
		if request.Method == route.httpMethod && params.Get("method") == route.apiMethod && params.Get("api") == route.api {
			route.handler(writer, request)
			return
		}
	}
	writer.WriteHeader(500)
}

func FakeServerHandler(account string, password string) *TestServer {
	t := &TestServer{account: account, password: password}
	t.AddRoute("GET", "SYNO.API.Auth", "login", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
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
	})
	return t
}
