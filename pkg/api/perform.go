package api

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

func GET[T any](client Client, queryTransformer func(query url.Values)) (*ResponseWrapper[T], error) {
	req, err := client.NewGET(queryTransformer)
	if err != nil {
		return nil, err
	}
	return PerformRequest[T](req)
}

func POST[T any](client Client, queryTransformer func(query url.Values)) (*ResponseWrapper[T], error) {
	req, err := client.NewPOST(queryTransformer)
	if err != nil {
		return nil, err
	}
	return PerformRequest[T](req)
}

func PerformRequest[T any](req *http.Request) (*ResponseWrapper[T], error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defaultLogger.Debug(string(body))
	return ParseResponse[T](bytes.NewBuffer(body))
}
