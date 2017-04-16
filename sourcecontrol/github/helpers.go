package github

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func newHttpClient() http.Client {
	return http.Client{}
}

func newHttpRequest(path string) (*http.Request, error) {
	apiUrl := GitHubApiHost + path

	fmt.Printf("URL: %s\n", apiUrl)

	request, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return request, err
	}
	request.Header.Add("User-Agent", "MyGoApp")

	return request, nil
}

func getAndUnmarshall(path string, out interface{}) error {
	client := newHttpClient()
	request, err := newHttpRequest(path)
	if err != nil {
		return err
	}

	response, err := client.Do(request)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, out)
	if err != nil {
		return err
	}

	return nil
}
