package api

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type client struct {
	BaseURL   string
	SessionID string
}

var defaultLogger *zap.Logger

func init() {
	defaultLogger = zap.NewNop()
}

// SetLogger overrides the default no-op zap.Logger with a user-provided logger.
// Use this for debug logging of http requests and responses.
func SetLogger(logger *zap.Logger) {
	defaultLogger = logger
}

type Client interface {
	Login(user, password string) (*ResponseWrapper[*LoginResponse], error)
	GetInfo() (*ResponseWrapper[*Info], error)
	NewGETRequest(queryTransformer func(query url.Values)) (*http.Request, error)
	NewPOSTRequest(queryTransformer func(query url.Values)) (*http.Request, error)
}

func NewClient(baseURL string) (Client, error) {
	c := &client{BaseURL: baseURL}
	return c, nil
}

func NewClientWithSessionID(baseURL, sessionID string) (Client, error) {
	return &client{BaseURL: baseURL, SessionID: sessionID}, nil
}

func PerformRequest[T any](client Client, queryTransformer func(query url.Values)) (*ResponseWrapper[T], error) {
	req, err := client.NewGETRequest(queryTransformer)
	if err != nil {
		return nil, err
	}
	return performRequest[T](req)
}

func performRequest[T any](req *http.Request) (*ResponseWrapper[T], error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defaultLogger.Debug(string(body))
	return ParseResponse[T](bytes.NewBuffer(body))
}

func PerformPOSTRequest[T any](client Client, queryTransformer func(query url.Values)) (*ResponseWrapper[T], error) {
	req, err := client.NewPOSTRequest(queryTransformer)
	if err != nil {
		return nil, err
	}
	return performRequest[T](req)
}
