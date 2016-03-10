package datas

import "time"

type Repository struct {
    ID int `json:"id"`
    Name string `json:"name"`
    FullName string `json:"full_name"`
    Owner User `json:"owner"`
    Private bool `json:"private"`
    HTMLURL string `json:"html_url"`
    Description string `json:"description"`
    Fork bool `json:"fork"`
    URL string `json:"url"`
    DeploymentsURL string `json:"deployments_url"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
    PushedAt *time.Time `json:"pushed_at"`
    CloneURL string `json:"clone_url"`
    SvnURL string `json:"svn_url"`
    Homepage string `json:"homepage"`
    Size int `json:"size"`
    StargazersCount int `json:"stargazers_count"`
    WatchersCount int `json:"watchers_count"`
    Language string `json:"language"`
    HasIssues bool `json:"has_issues"`
    HasDownloads bool `json:"has_downloads"`
    ForksCount int `json:"forks_count"`
    MirrorURL interface{} `json:"mirror_url"`
    OpenIssuesCount int `json:"open_issues_count"`
    Forks int `json:"forks"`
    OpenIssues int `json:"open_issues"`
    Watchers int `json:"watchers"`
    DefaultBranch string `json:"default_branch"`
    Score float64 `json:"score"`
    LanguageStats map[string]interface{} `json:"language_stats"`
}

type RepositoryBySize []Repository

func (a RepositoryBySize) Len() int           { return len(a) }
func (a RepositoryBySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a RepositoryBySize) Less(i, j int) bool { return a[i].Size > a[j].Size }
