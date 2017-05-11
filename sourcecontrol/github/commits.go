package github

import (
	"github.com/michaellockwood/github-repos/sourcecontrol"
	"time"
	"fmt"
)

const CommitPathFormat string = "/repos/%s/%s/commits?per_page=%v"

type author struct {
	Name string
	Date time.Time
}

type commit struct {
	URL string
	Author author
	Message string
}

type commitResponse struct {
	SHA string
	Commit commit
}

func (gateway GitHubGateway) GetCommits(owner string, repoName string, total int) ([]sourcecontrol.Commit, error) {
	commitResponses := []commitResponse{}
	err := getAndUnmarshall(fmt.Sprintf(CommitPathFormat, owner, repoName, total), &commitResponses)

	if err != nil {
		return nil, err
	}

	commits := make([]sourcecontrol.Commit, len(commitResponses))
	for i, c := range commitResponses {
		commits[i] = c.mapToDomain()
	}

	return commits, nil
}

func (commitResponse *commitResponse) mapToDomain() sourcecontrol.Commit {
	return sourcecontrol.Commit{
		Author: commitResponse.Commit.Author.Name,
		Date: commitResponse.Commit.Author.Date,
		Message: commitResponse.Commit.Message,
		SHA: commitResponse.SHA,
	}
}
