package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ParseResponse[T any](responseBody io.Reader) (*ResponseWrapper[T], error) {
	body, err := io.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}
	var jsonResponse *ResponseWrapper[T]
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}
	if !jsonResponse.Success {
		return nil, fmt.Errorf("failed request")
	}
	return jsonResponse, err
}

func (c *client) NewRequest(queryTransformer func(query url.Values)) (*http.Request, error) {
	return c.NewRequestWithoutAuth(func(query url.Values) {
		query.Add("_sid", c.SessionID)
		queryTransformer(query)
	})
}

func (c *client) NewRequestWithoutAuth(queryTransformer func(query url.Values)) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/webapi/entry.cgi", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	queryTransformer(query)
	req.URL.RawQuery = query.Encode()
	return req, nil
}
