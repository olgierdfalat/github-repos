package main

import (
	"fmt"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol/github"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol"
	"time"
	"os"
)

func main() {
	srcService := sourcecontrol.SourceControlService{github.GitHubGateway{}}

	repos, err := srcService.GetRepositoriesWithCommits("test", 5)

	if err != nil {
		exit(err)
	}

	for _, repo := range repos {
		fmt.Println("----------------------------------------------")
		fmt.Printf("Name: %s, Owner: %s, URL: %s, Created: %s, Last Push: %s\n", repo.Name, repo.Owner, repo.URL, repo.Created.Format(time.RFC3339), repo.LastPushed.Format(time.RFC3339))

		for _, commit := range repo.Commits {
			fmt.Printf("Author: %s, SHA: %s, Message: %s, Date: %s\n", commit.Author, commit.SHA, commit.Message, commit.Date.Format(time.RFC3339))
		}
		fmt.Println("---------------------------------------------")
	}
}

func exit(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("No error")
}
