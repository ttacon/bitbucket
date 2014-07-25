package bitbucket

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TeamService struct{}

type Team struct {
	Username    string `json:"username"`
	Kind        string `json:"kind"`
	Website     string `json:"website"`
	DisplayName string `json:"display_name"`
	Links       Links  `json:"links"`
	CreatedOn   string `json:"created_on"`
	Location    string `json:"location"`
}

func (t TeamService) GetTeamProfile(team string) (*Team, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/teams/%s", V2_URL, team),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Team
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type TeamMembers struct {
	PageLen int          `json:"pagelen"`
	Values  []TeamMember `json:"values"`
	Page    int          `json:"page"`
	Size    int          `json:"size"`
}

// TODO(ttacon): combine with owner/user?

type TeamMember struct {
	Username    string `json:"username"`
	Kind        string `json:"kind"`
	Website     string `json:"website"`
	DisplayName string `json:"display_name"`
	Links       Links  `json:"links"`
	CreatedOn   string `json:"created_on"`
	Location    string `json:"location"`
}

func (t TeamService) GetTeamMembers(team string) (*TeamMembers, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/teams/%s/members", V2_URL, team),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data TeamMembers
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
