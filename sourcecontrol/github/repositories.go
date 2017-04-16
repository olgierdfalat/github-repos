package github

import (
	"time"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const RepoSearchPathFormat string = "/search/repositories?q=%s&per_page=%v"

type owner struct {
	Login string
}

type repository struct {
	Name string
	Owner owner
	Html_url string
	Created_at time.Time
	Pushed_at time.Time
}

type searchResponse struct {
	Items []repository
}

func (gateway GitHubGateway) GetRepositories(query string, total int) ([]sourcecontrol.Repository, error) {
	client := newHttpClient()
	request, err := newHttpRequest(fmt.Sprintf(RepoSearchPathFormat, query, total))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse searchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.mapToDomain(), nil
}

func (response *searchResponse) mapToDomain() []sourcecontrol.Repository {
	repos := make([]sourcecontrol.Repository, len(response.Items))
	for i, item := range response.Items {
		repos[i] = sourcecontrol.Repository{
			Name: item.Name,
			Owner: item.Owner.Login,
			URL: item.Html_url,
			LastPushed: item.Created_at,
			Created: item.Pushed_at,
		}
	}

	return repos
}
