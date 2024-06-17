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
		"SYNO.API.Auth":         true,
		"SYNO.API.Info":         true,
		"SYNO.Docker.Container": true,
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
	analyze(har)
}

func analyze(har HAR) {
	count := 0
	for _, entry := range har.Log.Entries {
		u, err := url.Parse(entry.Request.Url)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if strings.HasPrefix(u.Path, "/webapi/entry.cgi") {
			if api(entry) != "" && !knownAPIs[api(entry)] {
				fmt.Println("new api?: ", api(entry))
				fmt.Println(entry.Request.PostData.Params)
				count++
				if count > 5 {
					break
				}
			}
		}
	}
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
