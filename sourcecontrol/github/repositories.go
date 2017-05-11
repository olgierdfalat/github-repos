package github

import (
	"time"
	"github.com/michaellockwood/github-repos/sourcecontrol"
	"fmt"
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
	var searchResponse searchResponse
	err := getAndUnmarshall(fmt.Sprintf(RepoSearchPathFormat, query, total), &searchResponse)
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
