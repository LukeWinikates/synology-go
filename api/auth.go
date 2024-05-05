package api

import (
	"net/http"
	"net/url"
)

func (c *client) Login() (*ResponseWrapper[*AuthResponse], error) {
	req, err := c.newRequest(func(query url.Values) {
		query.Add("api", "SYNO.API.Auth")
		query.Add("version", "6")
		query.Add("method", "login")
		query.Add("account", c.Account)
		query.Add("passwd", c.Password)
		query.Add("session", "FileStation")
		query.Add("format", "sid")
	})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	authResponse, err := ParseResponse[*AuthResponse](resp.Body)
	if err != nil {
		return nil, err
	}
	c.SessionID = authResponse.Data.Sid
	return authResponse, err
}
