package api

type Info map[string]APIDescription
type APIDescription struct {
	MaxVersion    int    `json:"maxVersion"`
	MinVersion    int    `json:"minVersion"`
	Path          string `json:"path"`
	RequestFormat string `json:"requestFormat"`
}

type AuthResponse struct {
	Did          string `json:"did"`
	IsPortalPort bool   `json:"is_portal_port"`
	Sid          string `json:"sid"`
}

type ResponseWrapper[T any] struct {
	Data    T    `json:"data"`
	Success bool `json:"success"`
}
