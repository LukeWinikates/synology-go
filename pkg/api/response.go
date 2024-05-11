package api

type Info map[string]Description
type Description struct {
	MaxVersion    int    `json:"maxVersion"`
	MinVersion    int    `json:"minVersion"`
	Path          string `json:"path"`
	RequestFormat string `json:"requestFormat"`
}

type LoginResponse struct {
	Did          string `json:"did"`
	IsPortalPort bool   `json:"is_portal_port"`
	Sid          string `json:"sid"`
}

type Error map[string]interface{}

type ResponseWrapper[T any] struct {
	Data    T     `json:"data"`
	Error   Error `json:"error"`
	Success bool  `json:"success"`
}

type ListResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
