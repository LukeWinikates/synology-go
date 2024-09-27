package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/LukeWinikates/synology-go/internal"
)

type mockRoute struct {
	api        string
	apiMethod  string
	httpMethod string
	handler    http.HandlerFunc
}

type TestServer struct {
	routes     []mockRoute
	Requests   []*http.Request
	HTTPServer *httptest.Server
	sessionID  string
}

func (t *TestServer) wrapAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		params := request.URL.Query()
		if request.Method == http.MethodPost {
			internal.Must(request.ParseForm())
			params = request.Form
		}

		if params.Get("_sid") == t.sessionID {
			handler(writer, request)

		} else {
			internal.Must(json.NewEncoder(writer).Encode(&ResponseWrapper[interface{}]{
				Data: nil,
				Error: Error{
					"code": 119,
				},
				Success: false,
			}))
		}
	}
}

func (t *TestServer) AddRoute(httpMethod, api, apiMethod string, handler http.HandlerFunc) {
	if api != "SYNO.API.Auth" {
		handler = t.wrapAuth(handler)
	}
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
func (t *TestServer) Session() string {
	return t.sessionID
}

func FakeServerHandler() *TestServer {
	t := &TestServer{
		sessionID: strconv.Itoa(rand.Int()),
	}
	return t
}
