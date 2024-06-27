package api

import (
	"net/http"

	"go.uber.org/zap"
)

type client struct {
	BaseURL    string
	Authorizer Authorizer
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
	NewGET(transformer ValueTransformer) (*http.Request, error)
	NewPOST(transformer ValueTransformer) (*http.Request, error)
}

func NewClient(baseURL string, authorizer Authorizer) Client {
	return &client{BaseURL: baseURL, Authorizer: authorizer}
}
