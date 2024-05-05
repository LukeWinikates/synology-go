package api

import "C"

type client struct {
	Account   string
	Password  string
	BaseURL   string
	SessionID string
}

type Client interface {
	Login() (*ResponseWrapper[*AuthResponse], error)
	GetInfo() (*ResponseWrapper[*Info], error)
}

func NewClient(baseURL, user, password string) (Client, error) {
	return &client{BaseURL: baseURL, Account: user, Password: password}, nil
}
