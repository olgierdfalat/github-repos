package main

import (
	"fmt"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol/github"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol"
)

func main() {
	srcService := sourcecontrol.SourceControlService{github.GitHubGateway{}}

	repos := srcService.GetRepositoriesWithCommits("test", 5)

	for _, repo := range repos {
		fmt.Println(repo.Name)
	}
}
