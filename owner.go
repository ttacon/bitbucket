package bitbucket

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
}

type Link struct {
	Href string `json:"href"`
}

type CloneLink struct {
	Href string `json:"href"`
	Name string `json:"name"`
}
