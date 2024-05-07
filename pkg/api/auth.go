package api

import (
	"net/http"
	"net/url"
)

func (c *client) Login() (*ResponseWrapper[*LoginResponse], error) {
	req, err := c.NewRequestWithoutAuth(func(query url.Values) {
		query.Add("api", "SYNO.API.Auth")
		query.Add("version", "6")
		query.Add("method", "login")
		query.Add("account", c.Account)
		query.Add("passwd", c.Password)
		query.Add("session", "FileStation")
		query.Add("format", "sid")
	})
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	authResponse, err := ParseResponse[*LoginResponse](resp.Body)
	if err != nil {
		return nil, err
	}
	c.SessionID = authResponse.Data.Sid
	return authResponse, err
}
