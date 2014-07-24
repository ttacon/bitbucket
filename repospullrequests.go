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

type Commit struct {
	Hash  string `json:"hash"`
	Links Links  `json:"links"`
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
	Description       string      `json:"description"`
	Links             Links       `json:"links"`
	Author            Owner       `json:"author"`
	CloseSourceBranch bool        `json:"close_source_branch"`
	Title             string      `json:"title"`
	Destination       Destination `json:"destination"`
	Reason            string      `json:"reason"`
	ClosedBy          *Owner      `json:"closed_by"`
	Source            Source      `json:"source"`
	State             string      `json:"state"`
	CreatedOn         string      `json:"created_on"`
	UpdatedOn         string      `json:"updated_on"`
	MergedCommit      *Commit     `json:"merge_commit"`
	Id                int         `json:"id"`
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
