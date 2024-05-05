package api

import (
	"net/http"
	"net/url"
)

func (c *client) GetInfo() (*ResponseWrapper[*Info], error) {
	req, err := c.newRequest(func(query url.Values) {
		query.Add("api", "SYNO.API.Info")
		query.Add("version", "1")
		query.Add("method", "query")
		query.Add("_sid", c.SessionID)
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ParseResponse[*Info](resp.Body)
}
