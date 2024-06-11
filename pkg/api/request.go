package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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
		return nil, fmt.Errorf("failed request: %s", jsonResponse.Error.ErrorCodeDescription())
	}
	return jsonResponse, err
}

func (c *client) NewGETRequest(queryTransformer func(query url.Values)) (*http.Request, error) {
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

func (c *client) NewPOSTRequest(queryTransformer func(query url.Values)) (*http.Request, error) {
	form := url.Values{}
	queryTransformer(form)
	form.Add("_sid", c.SessionID)

	req, err := http.NewRequest(
		"POST",
		c.BaseURL+"/webapi/entry.cgi/"+form.Get("api"),
		strings.NewReader(form.Encode()),
	)
	return req, err
}

func WrapQuote(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}
