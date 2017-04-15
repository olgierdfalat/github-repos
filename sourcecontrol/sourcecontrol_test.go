package sourcecontrol_test

import (
	"testing"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol"
	"github.com/golang/mock/gomock"
	"time"
	"bitbucket.org/michaellockwood/github-repos/sourcecontrol/mock_sourcecontrol"
)

func TestGetRepositoriesWithCommits_CommitsAreReturned(t *testing.T) {

	expectedTime := time.Date(2017, 1, 28, 12, 0, 0, 0, time.UTC)

	expectedRepos := []sourcecontrol.Repository{
		{"go", "golang", "www.go.com", expectedTime, expectedTime, nil},
		{"c#", "microsoft", "www.microsoft.com", expectedTime, expectedTime, nil},
	}

	expectedRepo1Commits := []sourcecontrol.Commit{
		{ "Google", "54544545", "Commit num 1", time.Now() },
		{ "Google", "gfdg43", "Commit num 2", time.Now() },
		{ "Google", "543ghxf", "Commit num 3", time.Now() },
		{ "Google", "868bc", "Commit num 4", time.Now() },
		{ "Google", "dfgdgf5", "Commit num 5", time.Now() },
	}

	expectedRepo2Commits := []sourcecontrol.Commit{
		{ "Microsoft", "xvxcvxbfb77756565", "Commit num 1", time.Now() },
		{ "Microsoft", "fdgdfg43334", "Commit num 2", time.Now() },
		{ "Microsoft", "df5555", "Commit num 3", time.Now() },
		{ "Microsoft", "bvbcb8787", "Commit num 4", time.Now() },
		{ "Microsoft", "cv4", "Commit num 5", time.Now() },
	}

	ctrl := gomock.NewController(t)
	gatewayMock := mock_sourcecontrol.NewMockSourceControlGateway(ctrl)
	gatewayMock.EXPECT().
		GetRepositories("*", 5).
		Return(expectedRepos)
	gatewayMock.EXPECT().
		GetCommits("golang", "go", 5).
		Return(expectedRepo1Commits)
	gatewayMock.EXPECT().
		GetCommits("microsoft", "c#", 5).
		Return(expectedRepo2Commits)
	service := sourcecontrol.SourceControlService{gatewayMock}

	actual := service.GetRepositoriesWithCommits("*", 5)

	if len(actual) != len(expectedRepos) {
		t.Fatalf("Expected: length %v, Actual: length %v", len(expectedRepos), len(actual))
	}

	RepoEquals(expectedRepos[0], actual[0], t)
	RepoEquals(expectedRepos[1], actual[1], t)

	CommitsEqual(expectedRepo1Commits, expectedRepos[0].Commits, t)
	CommitsEqual(expectedRepo2Commits, expectedRepos[1].Commits, t)
}

func RepoEquals(expectedRepo sourcecontrol.Repository, actualRepo sourcecontrol.Repository, t *testing.T) {
	if expectedRepo.Name != actualRepo.Name {
		t.Fatalf("Expected: Name %s, Actual: Name %s", expectedRepo.Name, actualRepo.Name)
	}

	if expectedRepo.Owner != actualRepo.Owner {
		t.Fatalf("Expected: Owner %s, Actual: Owner %s", expectedRepo.Owner, actualRepo.Owner)
	}

	if expectedRepo.URL != actualRepo.URL {
		t.Fatalf("Expected: URL %s, Actual: URL %s", expectedRepo.URL, actualRepo.URL)
	}

	if expectedRepo.Created != actualRepo.Created {
		t.Fatalf("Expected: Created %s, Actual: Created %s", expectedRepo.Created.Format(time.RFC3339), actualRepo.Created.Format(time.RFC3339))
	}

	if expectedRepo.LastPushed != actualRepo.LastPushed {
		t.Fatalf("Expected: LastPushed %s, Actual: LastPushed %s", expectedRepo.LastPushed.Format(time.RFC3339), actualRepo.LastPushed.Format(time.RFC3339))
	}
}

func CommitsEqual(expectedCommits []sourcecontrol.Commit, actualCommits []sourcecontrol.Commit, t *testing.T) {
	if len(expectedCommits) != len(actualCommits) {
		t.Fatalf("Expected: Commits length %v, Actual: %v", len(expectedCommits), len(actualCommits))
	}

	for i := range expectedCommits {
		expectedCommit := expectedCommits[i]
		actualCommit := actualCommits[i]

		if expectedCommit.Date != actualCommit.Date {
			t.Fatalf("Expected: Date %s, Actual: Date %s", expectedCommit.Date.Format(time.RFC3339), actualCommit.Date.Format(time.RFC3339))
		}

		if expectedCommit.Message != actualCommit.Message {
			t.Fatalf("Expected: Message %s, Actual: Message %s", expectedCommit.Message, actualCommit.Message)
		}

		if expectedCommit.SHA != actualCommit.SHA {
			t.Fatalf("Expected: SHA %s, Actual: SHA %s", expectedCommit.SHA, actualCommit.SHA)
		}

		if expectedCommit.Author != actualCommit.Author {
			t.Fatalf("Expected: Author %s, Actual: Author %s", expectedCommit.Author, actualCommit.Author)
		}
	}
}