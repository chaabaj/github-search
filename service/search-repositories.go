/*
The MIT License (MIT)

Copyright (c) 2016 Chaabane Jalal

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package service

import (
       "fmt"
       "os"
       "errors"
       "encoding/json"
       "sort"
       "github.com/chaabaj/github-search/service/api"
       "github.com/chaabaj/github-search/datas"
)

// Represent data structure GitHub API Search response
type searchResult struct {
	TotalCount int `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items []datas.Repository `json:"items"`
}

// GitHub API definition with a authorization token
var githubApi = api.New("https://api.github.com", os.Getenv("GITHUB_AUTH_TOKEN"))

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

    // Dipatch nbRepositories calls to block of chunkSize
    // 1 go routine handle chunkSize of calls to getRepositoryLanguages
    // Errors are propageted using the channel reqChan
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

// Search GitHub repositories by name
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
    if len(result.Items) > 0 {
        return resolveRepositoryLanguage(result.Items)
    }
    return nil, errors.New("No Results")
}
