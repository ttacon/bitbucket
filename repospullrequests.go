package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PullRequests struct {
	PageLen int           `json:"page_len"`
	Next    string        `json:"next"`
	Values  []PullRequest `json:"values"`
	Page    int           `json:"page"`
	Size    int           `json:"size"`
}

type Destination struct {
	Commit     Commit     `json:"commit"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

type Source struct {
	Commit     Commit     `json:"commit"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type PullRequest struct {
	Description       string        `json:"description"`
	Links             Links         `json:"links"`
	Author            Owner         `json:"author"`
	CloseSourceBranch bool          `json:"close_source_branch"`
	Title             string        `json:"title"`
	Destination       Destination   `json:"destination"`
	Reason            string        `json:"reason"`
	ClosedBy          *Owner        `json:"closed_by"`
	Source            Source        `json:"source"`
	State             string        `json:"state"`
	CreatedOn         string        `json:"created_on"`
	UpdatedOn         string        `json:"updated_on"`
	MergedCommit      *Commit       `json:"merge_commit"`
	Id                int           `json:"id"`
	Reviewers         []Owner       `json:"reviewers"`
	Participants      []Participant `json:"participants"`
}

type Participant struct {
	User     Owner  `json:"user"`
	Role     string `json:"role"`
	Approved bool   `json:"approved"`
}

func (r RepositoryService) GetPullRequests(owner, repo string, states ...string) (*PullRequests, error) {
	// TODO(ttacon): validate states are members of [OPEN, MERGED, DECLINED]
	var queryString string
	if len(states) > 0 {
		queryString = "?state=" + strings.Join(states, ",")
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s/pullrequests%s",
			V2_URL,
			owner,
			repo,
			queryString),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data PullRequests
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r RepositoryService) GetPullRequest(owner, repo, requestId string) (*PullRequest, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s/pullrequests/%s",
			V2_URL,
			owner,
			repo,
			requestId),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data PullRequest
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type Commit struct {
	Hash       string     `json:"hash"`
	Links      Links      `json:"links"`
	Repository Repository `json:"repository"`
	Author     Author     `json:"author"`
	Parents    []Commit   `json:"parents"`
	Date       string     `json:"date"`
	Message    string     `json:"message"`
}

type Author struct {
	Raw  string `json:"raw"`
	User Owner  `json:"user"`
}

type Commits struct {
	PageLen int      `json:"page_len"`
	Next    string   `json:"next"`
	Values  []Commit `json:"values"`
	Page    int      `json:"page"`
	Size    int      `json:"size"`
}

func (r RepositoryService) GetPullRequestCommits(owner, repo, requestId string) (*Commits, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s/pullrequests/%s/commits",
			V2_URL,
			owner,
			repo,
			requestId),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Commits
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
