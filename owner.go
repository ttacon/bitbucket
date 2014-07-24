package bitbucket

// TODO(ttacon): change name of Owner to user or something similar

type Owner struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Links       Links  `json:"links"`
}

type Links struct {
	Self         Link        `json:"self,omitempty"`
	Avatar       Link        `json:"avatar,omitempty"`
	Watchers     Link        `json:"watchers"`
	Commits      Link        `json:"commits"`
	HTML         Link        `json:"html"`
	Forks        Link        `json:"forks"`
	Clone        []CloneLink `json:"clone"`
	PullRequests Link        `json:"pullrequests"`

	// Links for pull requests
	Comments Link `json:"comments"`
	Patch    Link `json:"patch"`
	Merge    Link `json:"merge"`
	Activity Link `json:"activity"`
	Diff     Link `json:"diff"`
	Approve  Link `json:"approve"`
	Decline  Link `json:"decline"`

	// Links for teams
	Repositories Link `json:"repositories"`
	Followers    Link `json:"followers"`
	Members      Link `json:"members"`
	Following    Link `json:"following"`
}

type Link struct {
	Href string `json:"href"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}
