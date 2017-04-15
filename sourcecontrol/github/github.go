package github

import "bitbucket.org/michaellockwood/github-repos/sourcecontrol"

type GitHubGateway struct {
}

func (gateway GitHubGateway) GetRepositories(query string, total int) []sourcecontrol.Repository {
	return make([]sourcecontrol.Repository, 0)
}

func (gateway GitHubGateway) GetCommits(owner string, repoName string, total int) []sourcecontrol.Commit {
	return make([]sourcecontrol.Commit, 0)
}