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
	GetRepositories(query string, total int) ([]Repository, error)
	GetCommits(owner string, repoName string, total int) ([]Commit, error)
}

type SourceControlService struct {
	Gateway SourceControlGateway
}

func (sourceControlService *SourceControlService) GetRepositoriesWithCommits(query string, total int) ([]Repository, error) {
	repos, err := sourceControlService.Gateway.GetRepositories(query, total)
	if err != nil {
		return nil, err
	}
	for i := range repos {
		repo := &repos[i]
		repo.Commits, err = sourceControlService.Gateway.GetCommits(repo.Owner, repo.Name, total)
		if err != nil {
			return nil, err
		}
	}

	return repos, nil
}