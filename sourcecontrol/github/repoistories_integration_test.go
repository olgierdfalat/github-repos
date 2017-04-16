package github_test

import (
	"testing"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol/github"
	"time"
)

func TestGetRepositories_ReturnsCorrectAmountOfRepositories(t *testing.T) {
	cases := []struct {
		count int
	}{
		{5},
		{10},
	}

	for _, c := range cases {
		gateway := github.GitHubGateway{}
		repos, err := gateway.GetRepositories("test", c.count)

		if err != nil {
			t.Fatalf("GetRepositories failed with error: %v", err)
		}

		if len(repos) != c.count {
			t.Fatalf("Expected repo count: %v, Actual: %v", c.count, len(repos))
		}
	}
}

func TestGetRepositories_ReturnsDomainRepository(t *testing.T) {
	gateway := github.GitHubGateway{}
	repos, err := gateway.GetRepositories("test", 5)

	if err != nil {
		t.Fatalf("GetRepositories failed with error: %v", err)
	}

	for _, r := range repos {
		if r.Owner == "" {
			t.Fatal("Owner was empty.")
		}

		if r.Name == "" {
			t.Fatal("Name was empty.")
		}

		if r.URL == "" {
			t.Fatal("URL was empty.")
		}

		zeroTime := time.Time{}
		if r.LastPushed == zeroTime {
			t.Fatal("LastPushed was zero time.")
		}

		if r.Created == zeroTime {
			t.Fatal("Created was zero time.")
		}
	}
}