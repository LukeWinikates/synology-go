package api

import (
	"net/http"
	"net/url"
)

type client struct {
	BaseURL   string
	SessionID string
}

type Client interface {
	Login(user, password string) (*ResponseWrapper[*LoginResponse], error)
	GetInfo() (*ResponseWrapper[*Info], error)
	NewRequest(queryTransformer func(query url.Values)) (*http.Request, error)
}

func NewClient(baseURL string) (Client, error) {
	c := &client{BaseURL: baseURL}
	return c, nil
}

func NewClientWithSessionID(baseURL, sessionID string) (Client, error) {
	return &client{BaseURL: baseURL, SessionID: sessionID}, nil
}

func PerformRequest[T any](client Client, queryTransformer func(query url.Values)) (*ResponseWrapper[T], error) {
	req, err := client.NewRequest(queryTransformer)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ParseResponse[T](resp.Body)
}
