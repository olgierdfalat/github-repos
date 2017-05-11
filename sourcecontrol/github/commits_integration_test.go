package github_test

import (
	"testing"
	"github.com/michaellockwood/github-repos/sourcecontrol/github"
	"time"
)

func TestGitHubGateway_GetCommits_ReturnsCorrectAmountOfCommits(t *testing.T) {
	cases := []struct {
		count int
	}{
		{ 10 },
		{ 5 },
	}

	for _, c := range cases {
		gateway := github.GitHubGateway{}
		commits, err := gateway.GetCommits("golang", "go", c.count)

		if err != nil {
			t.Fatalf("GetRepositories failed with error: %v", err)
		}

		if len(commits) != c.count {
			t.Fatalf("Expected repo count: %v, Actual: %v", c.count, len(commits))
		}
	}
}

func TestGitHubGateway_GetCommits_ReturnsDomainCommits(t *testing.T) {
	gateway := github.GitHubGateway{}
	commits, err := gateway.GetCommits("golang", "go", 5)

	if err != nil {
		t.Fatalf("GetRepositories failed with error: %v", err)
	}

	for _, c := range commits {
		if c.Author == "" {
			t.Fatal("Author was empty.")
		}

		if c.SHA == "" {
			t.Fatal("SHA was empty.")
		}

		if c.Message == "" {
			t.Fatal("Message was empty.")
		}

		zeroTime := time.Time{}
		if c.Date == zeroTime {
			t.Fatal("Date was zero time.")
		}
	}
}
