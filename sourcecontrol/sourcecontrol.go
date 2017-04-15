package sourcecontrol

import "time"

type Repository struct {
	Name string
	Owner string
	URL string
	LastPushed time.Time
	Created time.Time
}

type Commit struct {
	Author string
	SHA string
	Message string
	Date time.Time
}

type SourceControlGateway interface {
	GetRepositories(query string, total int) Repository[]
	GetCommits(owner string, repoName string, total int) Commit[]
}
