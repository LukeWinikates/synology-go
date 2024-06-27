package auth

import (
	"net/url"

	"github.com/LukeWinikates/synology-go/pkg/api"
)

type PasswordLoginClient struct {
	apiClient api.Client
	baseURL   string
	sessionID string
}

func NewPasswordLoginClient(baseURL string) *PasswordLoginClient {
	client := api.NewClient(baseURL, api.NoopAuthorizer{})
	return &PasswordLoginClient{apiClient: client, baseURL: baseURL}
}

type LoginResponse struct {
	Did          string `json:"did"`
	IsPortalPort bool   `json:"is_portal_port"`
	Sid          string `json:"sid"`
}

func (a *PasswordLoginClient) Login(account, password string) (*api.ResponseWrapper[*LoginResponse], error) {
	authResponse, err := api.GET[*LoginResponse](a.apiClient, func(query url.Values) {
		query.Add("api", "SYNO.API.Auth")
		query.Add("version", "6")
		query.Add("method", "login")
		query.Add("account", account)
		query.Add("passwd", password)
		query.Add("session", "FileStation")
		query.Add("format", "sid")
	})
	if err != nil {
		return nil, err
	}
	a.sessionID = authResponse.Data.Sid
	return authResponse, err
}

func (a *PasswordLoginClient) NewAPIClient() api.Client {
	return api.NewClient(a.baseURL, NewSessionAuthorizer(a.sessionID))
}
