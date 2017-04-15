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
		t.Errorf("Expected: length %b, Actual: length %b", len(expectedRepos), len(actual))
	}

	RepoEquals(expectedRepos[0], actual[0], t)
	RepoEquals(expectedRepos[1], actual[1], t)
}

func RepoEquals(expectedRepo sourcecontrol.Repository, actualRepo sourcecontrol.Repository, t *testing.T) {
	if expectedRepo.Name != actualRepo.Name {
		t.Errorf("Expected: Name %s, Actual: Name %s", expectedRepo.Name, actualRepo.Name)
	}

	if expectedRepo.Owner != actualRepo.Owner {
		t.Errorf("Expected: Owner %s, Actual: Owner %s", expectedRepo.Owner, actualRepo.Owner)
	}

	if expectedRepo.URL != actualRepo.URL {
		t.Errorf("Expected: URL %s, Actual: URL %s", expectedRepo.URL, actualRepo.URL)
	}

	if expectedRepo.Created != actualRepo.Created {
		t.Errorf("Expected: Created %s, Actual: Created %s", expectedRepo.Created.Format(time.RFC3339), actualRepo.Created.Format(time.RFC3339))
	}

	if expectedRepo.LastPushed != actualRepo.LastPushed {
		t.Errorf("Expected: LastPushed %s, Actual: LastPushed %s", expectedRepo.LastPushed.Format(time.RFC3339), actualRepo.LastPushed.Format(time.RFC3339))
	}

	if len(expectedRepo.Commits) != len(actualRepo.Commits) {
		t.Errorf("Expected: Commits length %b, Actual: %b", len(expectedRepo.Commits), len(actualRepo.Commits))
	}

	for i := range expectedRepo.Commits {
		expectedCommit := expectedRepo.Commits[i]
		actualCommit := actualRepo.Commits[i]
		CommitEquals(expectedCommit, actualCommit, t)
	}
}

func CommitEquals(expectedCommit sourcecontrol.Commit, actualCommit sourcecontrol.Commit, t *testing.T) {
	if expectedCommit.Date != actualCommit.Date {
		t.Errorf("Expected: Date %s, Actual: Date %s", expectedCommit.Date.Format(time.RFC3339), actualCommit.Date.Format(time.RFC3339))
	}

	if expectedCommit.Message != actualCommit.Message {
		t.Errorf("Expected: Message %s, Actual: Message %s", expectedCommit.Message, actualCommit.Message)
	}

	if expectedCommit.SHA != actualCommit.SHA {
		t.Errorf("Expected: SHA %s, Actual: SHA %s", expectedCommit.SHA, actualCommit.SHA)
	}

	if expectedCommit.Author != actualCommit.Author {
		t.Errorf("Expected: Author %s, Actual: Author %s", expectedCommit.Author, actualCommit.Author)
	}
}