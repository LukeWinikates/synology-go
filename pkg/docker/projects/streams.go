package projects

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
)

func (pc *client) BuildStream(id string, lineReader func(s string)) error {
	return pc.streamingRequest(id, "build_stream", lineReader)
}

func (pc *client) streamingRequest(id string, method string, lineReader func(s string)) error {
	req, err := pc.apiClient.NewRequest(func(query url.Values) {
		query.Add("api", "SYNO.Docker.Project")
		query.Add("version", "1")
		query.Add("method", method)
		query.Add("id", fmt.Sprintf(`"%s"`, id))
	})
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		lineReader(scanner.Text())
	}
	return err
}
func (pc *client) CleanStream(id string, lineReader func(s string)) error {
	return pc.streamingRequest(id, "clean_stream", lineReader)
}

func (pc *client) Start(id string, lineReader func(s string)) error {
	return pc.streamingRequest(id, "start_stream", lineReader)
}

func (pc *client) Stop(id string, lineReader func(s string)) error {
	return pc.streamingRequest(id, "stop_stream", lineReader)
}
