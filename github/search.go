package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// APIBaseEndpoint contains base url to Github API
	APIBaseEndpoint = "https://api.github.com"

	// APISearchEndpoint contains url path to search endpoint
	APISearchEndpoint = "/search/repositories"
)

type SearchResult struct {
	TotalCount        int                 `json:"total_count"`
	IncompleteResults bool                `json:"incomplete_results"`
	Items             []SearchResultItems `json:"items"`
}

type SearchResultItems struct {
	Id              int       `json:"id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Owner           Owner     `json:"owner"`
	Private         bool      `json:"private"`
	HtmlUrl         string    `json:"html_url"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	Url             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	Homepage        string    `json:"homepage"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	ForksCount      int       `json:"forks_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	MasterBranch    string    `json:"master_branch"`
	DefaultBranch   string    `json:"default_branch"`
	Score           float32   `json:"score"`
}

func (i SearchResultItems) PackageName() string {
	return fmt.Sprintf("github.com/%s", strings.ToLower(i.FullName))
}

type Owner struct {
	Id                int    `json:"id"`
	Login             string `json:"login"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
}

func Search(query string) (*SearchResult, error) {
	fullQuery := fmt.Sprintf("%s language:go", query)
	params := url.Values{}
	params.Add("q", fullQuery)
	params.Add("sort", "stars")
	queryUrl := fmt.Sprintf("%s%s?%s", APIBaseEndpoint, APISearchEndpoint, params.Encode())
	response, err := http.Get(queryUrl)
	if err != nil {
		return nil, fmt.Errorf("http get error: %s", err)
	}
	defer response.Body.Close()
	result := new(SearchResult)
	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return nil, fmt.Errorf("json decode err: %s", err)
	}
	return result, nil
}
