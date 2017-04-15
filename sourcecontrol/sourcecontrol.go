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

type sourceControlGateway interface {
	getRepositories(query string, total int) Repository[]
	getCommits(owner string, repoName string, total int) Commit[]
}

type SourceControlService interface {
	GetRepositoriesWithCommits(query string, total int) Repository[]
}
