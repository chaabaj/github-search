package api

import (
       "net/http"
       "io/ioutil"
       "github.com/jalalc-github-search/utils"
       "errors"
)

// Basic representation of an Api with his base url
type Api struct {
     baseUrl string
     authToken string
}

// Create a new instance of Api
func New(baseUrl string, authToken string) *Api {
     return &Api{baseUrl : baseUrl, authToken : authToken}
}

// Call a get method on the service with the get parameters
// It return the response as an array of bytes
// return an error if something wrong occur
func (api *Api) Get(name string, params map[string]string) ([]byte, error) {
    req, err := http.NewRequest("GET", api.baseUrl + "/" + name, nil)
    client := &http.Client{}

    if err != nil {
       return nil, err
    }
    query := req.URL.Query()
    for key, val := range params {
    	query.Add(key, val)
    }
    req.URL.RawQuery = query.Encode()
    req.Header.Add("Authorization", "token " + api.authToken)
    utils.Log.Println("Sending request at : ", req.URL.String())
    resp, err := client.Do(req)
    if err != nil {
       return nil, err
    }
    if resp.StatusCode >= 400 {
        return nil, errors.New(http.StatusText(resp.StatusCode))
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
       return nil, err
    }
    return body, nil
}

// Only Get is implemented for Api.go but we can imagine some other method
