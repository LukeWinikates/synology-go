package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ParseResponse[T any](responseBody io.ReadCloser) (*ResponseWrapper[T], error) {
	defer responseBody.Close()
	var jsonResponse *ResponseWrapper[T]
	err := json.NewDecoder(responseBody).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}
	if !jsonResponse.Success {
		return nil, fmt.Errorf("failed request")
	}
	return jsonResponse, err
}

func (c *client) newRequest(queryTransformer func(query url.Values)) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/webapi/entry.cgi", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	queryTransformer(query)
	req.URL.RawQuery = query.Encode()
	return req, nil
}
