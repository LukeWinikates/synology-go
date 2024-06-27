package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

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
