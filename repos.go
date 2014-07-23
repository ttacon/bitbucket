package bitbucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RepositoryService struct {
}

type Repository struct {
	HasWiki     bool             `json:"has_wiki,omitempty"`
	Name        string           `json:"name,omitempty"`
	SCM         string           `json:"scm,omitempty"`
	ForkPolicy  string           `json:"fork_policy,omitempty"` //should make this some sort of enum?
	FullName    string           `json:"full_name"`
	Owner       Owner            `json:"owner"`
	Size        int              `json:"size"`
	IsPrivate   bool             `json:"is_private"`
	Description string           `json:"description,omitempty"`
	Links       Links            `json:"links"`
	HasIssues   bool             `json:"has_issues"`
	Language    string           `json:"language"`
	CreatedOn   string           `json:"created_on"`
	UpdatedOn   string           `json:"updated_on"`
	Parent      ParentRepository `json:"parent"`
}

type ParentRepository struct {
	Links    Links  `json:"links"`
	FullName string `json:"full_name"`
	Name     string `json:"name"`
}

func (r RepositoryService) GetRepo(owner, repo string) (*Repository, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s", V2_URL, owner, repo),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Repository
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type Forks struct {
	PageLen int    `json:"pagelen"`
	Values  []Fork `json:"values"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

type Fork struct {
	SCM         string           `json:"scm"`
	HasWiki     bool             `json:"has_wiki"`
	Description string           `json:"description"`
	Links       Links            `json:"links"`
	ForkPolicy  string           `json:"fork_policy"`
	Language    string           `json:"language"`
	CreatedOn   string           `json:"created_on"`
	Parent      ParentRepository `json:"parent"`
	FullName    string           `json:"full_name"`
	HasIssues   bool             `json:"has_issues"`
	Owner       Owner            `json:"owner"`
	UpdatedOn   string           `json:"updated_on"`
	Size        int              `json:"size"`
	IsPrivate   bool             `json:"is_private"`
}

func (r RepositoryService) GetForks(owner, repo string) (*Forks, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s/forks", V2_URL, owner, repo),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Forks
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// https://confluence.atlassian.com/display/BITBUCKET/repository+Resource#repositoryResource-GETalistofwatchers
//
type Watchers struct {
	Page    int       `json:"page,omitempty"`
	Size    int       `json:"size,omitempty"`
	PageLen int       `json:"pagelen,omitempty"`
	Next    string    `json:"next,omitempty"`
	Values  []Watcher `json:"values,omitempty"`
}

type Watcher struct {
	Username    string `json:"username,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Links       Links  `json:"links,omitempty"`
}

func (r RepositoryService) GetWatchers(owner, repo string) (*Watchers, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/repositories/%s/%s/watchers", V2_URL, owner, repo),
		nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Watchers
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp)

	return &data, nil
}

// TODO(ttacon): ughness everywhere, As Lorde says, "I'm getting ready to throw my hands up in the air"
func (r RepositoryService) CreateRepository(owner string, repo *Repository) (*Watchers, error) {
	if repo == nil {
		return nil, fmt.Errorf("repo cannot be nil")
	}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(repo); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("%s/repositories/%s/%s", V2_URL, owner, repo.Name),
		buf)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", *resp)

	//	var data map[string]interface{}
	bs, err := ioutil.ReadAll(resp.Body)
	fmt.Println("err: ", err)
	fmt.Println("body: ", string(bs))
	/*
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, err
		}
		fmt.Println(data)
	*/

	return nil, nil
}

// TODO(ttacon): add auth
func (r RepositoryService) DeleteRepository(owner, repo string) error {
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("%s/repositories/%s/%s", V2_URL, owner, repo),
		nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}
