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

type ResponseWrapper[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
}

type ListResponse struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}
