package github

import (
	"net/http"
	"fmt"
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
