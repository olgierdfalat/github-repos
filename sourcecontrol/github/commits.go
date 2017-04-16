package github

import (
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol"
	"time"
)

const CommitPathFormat string = "/repos/%s/%s/commits?per_page=%v"

func (gateway GitHubGateway) GetCommits(owner string, repoName string, total int) ([]sourcecontrol.Commit, error) {
	return []sourcecontrol.Commit {
		{ "Google", "54544545", "Commit num 1", time.Now() },
		{ "Google", "gfdg43", "Commit num 2", time.Now() },
		{ "Google", "543ghxf", "Commit num 3", time.Now() },
		{ "Google", "868bc", "Commit num 4", time.Now() },
		{ "Google", "dfgdgf5", "Commit num 5", time.Now() },
	}, nil
}
