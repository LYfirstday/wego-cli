package utils

type GithubResponseFile struct {
	Name         string `json:"name"`
	Content      string `json:"content"`
	Path         string `json:"path"`
	Sha          string `json:"sha"`
	Size         int64  `json:"size"`
	Url          string `json:"url"`
	Html_url     string `json:"html_url"`
	Git_url      string `json:"git_url"`
	Download_url string `json:"download_url"`
	Type         string `json:"type"`
	Encoding     string `json:"encoding"`
	Links        struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		Html string `json:"html"`
	} `json:"_links"`
}
