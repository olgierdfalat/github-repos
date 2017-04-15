package sourcecontrol

import "time"

type Repository struct {
	Name string
	Owner string
	URL string
	LastPushed time.Time
	Created time.Time
	Commits []Commit
}

type Commit struct {
	Author string
	SHA string
	Message string
	Date time.Time
}

type SourceControlGateway interface {
	GetRepositories(query string, total int) []Repository
	GetCommits(owner string, repoName string, total int) []Commit
}

type SourceControlService struct {
	Gateway SourceControlGateway
}

func (sourceControlService *SourceControlService) GetRepositoriesWithCommits(query string, total int) []Repository {
	repos := sourceControlService.Gateway.GetRepositories(query, total)
	for _, repo := range repos {
		repo.Commits = sourceControlService.Gateway.GetCommits(repo.Owner, repo.Name, total)
	}

	return repos
}