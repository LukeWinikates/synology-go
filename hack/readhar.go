package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"
)

var knownAPIs map[string]bool

func init() {
	knownAPIs = map[string]bool{
		"SYNO.API.Auth":             true,
		"SYNO.API.Info":             true,
		"SYNO.Docker.Container":     true,
		"SYNO.Docker.Container.Log": true,
	}
}

type Content struct {
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
	Text     string `json:"text"`
}

type Response struct {
	Status      int           `json:"status"`
	StatusText  string        `json:"statusText"`
	HttpVersion string        `json:"httpVersion"`
	Headers     []interface{} `json:"headers"`
	Cookies     []interface{} `json:"cookies"`
	Content     Content       `json:"content"`
	RedirectURL string        `json:"redirectURL"`
	HeadersSize int           `json:"headersSize"`
	BodySize    int           `json:"bodySize"`
}

type NameValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Request struct {
	BodySize    int         `json:"bodySize"`
	Method      string      `json:"method"`
	Url         string      `json:"url"`
	HttpVersion string      `json:"httpVersion"`
	PostData    PostData    `json:"postData"`
	Headers     []NameValue `json:"headers"`
	Cookies     []NameValue `json:"cookies"`
	QueryString []NameValue `json:"queryString"`
	HeadersSize int         `json:"headersSize"`
}

type PostData struct {
	MimeType string      `json:"mimeType"`
	Params   []NameValue `json:"params"`
	Text     string      `json:"text"`
}

type Entry struct {
	StartedDateTime time.Time `json:"startedDateTime"`
	Request         Request   `json:"request"`
	Response        Response  `json:"response"`
	Time            int       `json:"time"`
	Pageref         string    `json:"pageref"`
}

type Log struct {
	Entries []Entry `json:"entries"`
}

type HAR struct {
	Log Log `json:"log"`
}

func main() {
	file := os.Args[1]
	reader, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	var har HAR
	err = json.NewDecoder(reader).Decode(&har)
	if err != nil {
		panic(err)
	}
	report(analyze(har))
}

func report(APIs *DiscoveredAPIs) {
	for api, methodMap := range APIs.items {
		fmt.Println(api)
		for method, discoveredAPIS := range methodMap {
			fmt.Println(method)
			for _, a := range discoveredAPIS {
				fmt.Println(a.params)
			}
		}
	}
}

type DiscoveredAPI struct {
	API     string
	Method  string
	Version string
	params  []NameValue
}

type DiscoveredAPIs struct {
	items map[string]map[string][]DiscoveredAPI
}

func (d *DiscoveredAPIs) Add(api, method, version string, params []NameValue) {
	if _, ok := d.items[api][method]; ok {
		return
	}

	item, ok := d.items[api]
	if !ok {
		d.items[api] = map[string][]DiscoveredAPI{}
		item = d.items[api]
	}

	item[method] = append(item[method], DiscoveredAPI{
		API:     api,
		Method:  method,
		Version: version,
		params:  params,
	})
}

func analyze(har HAR) *DiscoveredAPIs {
	result := &DiscoveredAPIs{items: map[string]map[string][]DiscoveredAPI{}}
	for _, entry := range har.Log.Entries {
		u, err := url.Parse(entry.Request.Url)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if strings.HasPrefix(u.Path, "/webapi/entry.cgi") {
			if api(entry) != "" {
				result.Add(api(entry), method(entry), version(entry), rest(entry))
			}
		}
	}
	return result
}

func method(entry Entry) string {
	return get("method", entry)
}
func version(entry Entry) string {
	return get("version", entry)
}
func rest(entry Entry) []NameValue {
	var result []NameValue
	for _, param := range entry.Request.PostData.Params {
		if param.Name != "method" && param.Name != "version" && param.Name != "api" {
			result = append(result, param)
		}
	}
	for _, query := range entry.Request.QueryString {
		if query.Name != "method" && query.Name != "version" && query.Name != "api" {
			result = append(result, query)
		}
	}
	return result
}

func get(key string, entry Entry) string {
	for _, param := range entry.Request.PostData.Params {
		if param.Name == key {
			return param.Value
		}
	}
	for _, query := range entry.Request.QueryString {
		if query.Name == key {
			return query.Value
		}
	}
	return ""
}

func api(entry Entry) string {
	for _, param := range entry.Request.PostData.Params {
		if param.Name == "api" {
			return param.Value
		}
	}
	for _, query := range entry.Request.QueryString {
		if query.Name == "api" {
			return query.Value
		}
	}
	return ""
}
