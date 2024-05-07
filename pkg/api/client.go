package api

import (
	"net/http"
	"net/url"
)

type client struct {
	Account   string
	Password  string
	BaseURL   string
	SessionID string
}

type Client interface {
	Login() (*ResponseWrapper[*LoginResponse], error)
	GetInfo() (*ResponseWrapper[*Info], error)
	NewRequest(queryTransformer func(query url.Values)) (*http.Request, error)
}

func NewClient(baseURL, user, password string) (Client, error) {
	return &client{BaseURL: baseURL, Account: user, Password: password}, nil
}
