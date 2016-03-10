package service

import (
       "fmt"
       "encoding/json"
       "sort"
       "github.com/chaabaj/github-search/service/api"
       "github.com/chaabaj/github-search/datas"
)

type searchResult struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []datas.Repository `json:"items"`
}

// github api definition with a valid auth token
// It provide only access to public repository
var githubApi = api.New("https://api.github.com", "token bf107c4ecf091b0f09f5eb3e4fb77432589f5b21")

// Try to get the languages used in the repository
func getRepositoryLanguages(repo *datas.Repository) (map[string]interface{}, error) {
    service := fmt.Sprintf("repos/%s/%s/languages", repo.Owner.Login, repo.Name)
    body, err := githubApi.Get(service, map[string]string{})
    var langStats map[string]interface{}

    if err != nil {
        return nil, err
    } else if err := json.Unmarshal(body, &langStats); err != nil {
        return nil, err
    }
    return langStats, nil
}

// Try to resolve the languages of each repository in the repository array
// If it succeed it return the repositories that are updated with theirs languages
func resolveRepositoryLanguage(repositories []datas.Repository) ([]datas.Repository, error) {
    reqChan := make(chan error)
    nbRepositories:= len(repositories)
    chunkSize := nbRepositories / 10

    for i := 0; i < nbRepositories; i += chunkSize {
        go func(start int, end int) {
            for j := start; j < end && j < nbRepositories; j++ {
                stats, err := getRepositoryLanguages(&repositories[j])

                if err != nil {
                    reqChan <- err
                    return
                } else {
                    repositories[j].LanguageStats = stats
                }
            }
            reqChan <- nil
        }(i, i + chunkSize)
    }

    remaining := nbRepositories
    for {
        select {
        case err := <- reqChan:
            if err != nil {
                return nil, err
            }
            remaining -= chunkSize
            if remaining <= 0 {
                sort.Sort(datas.RepositoryBySize(repositories))
                return repositories, nil
            }
        }
    }
}

// Search repositories by name
// It return the repositories sorted by size
func SearchRepositories(name string) ([]datas.Repository, error) {
     var result searchResult

     params := map[string]string {
     	"q" : name + " in:name",
	    "type" : "repositories",
        "page" : "1",
        "per_page" : "100",
        "sort" : "stars",
        "order" : "desc",
     }
     body, err := githubApi.Get("search/repositories", params)
     if err != nil {
     	return nil, err
     }
     if err := json.Unmarshal(body, &result); err != nil {
     	return nil, err
     }
     return resolveRepositoryLanguage(result.Items)
}
