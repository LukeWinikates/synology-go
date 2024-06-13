package api

import (
	"net/http"
	"net/url"
)

func (c *client) GetInfo() (*ResponseWrapper[*Info], error) {
	req, err := c.NewGETRequest(func(query url.Values) {
		query.Add("api", "SYNO.API.Info")
		query.Add("version", "1")
		query.Add("method", "query")
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ParseResponse[*Info](resp.Body)
}
