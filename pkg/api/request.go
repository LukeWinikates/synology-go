package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ValueTransformer = func(values url.Values)

func WrapQuote(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func (c *client) NewGET(queryTransformer func(query url.Values)) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, c.BaseURL+"/webapi/entry.cgi", nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	queryTransformer(query)
	c.Authorizer.Apply(query)
	req.URL.RawQuery = query.Encode()
	return req, nil
}

func (c *client) NewPOST(formTransformer ValueTransformer) (*http.Request, error) {
	form := url.Values{}
	formTransformer(form)
	c.Authorizer.Apply(form)
	req, err := http.NewRequest(
		http.MethodPost,
		c.BaseURL+"/webapi/entry.cgi/"+form.Get("api"),
		strings.NewReader(form.Encode()),
	)
	return req, err
}
