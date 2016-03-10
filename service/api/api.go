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

package api

import (
       "net/http"
       "io/ioutil"
       "github.com/chaabaj/github-search/utils"
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
    query.Add("access_token", api.authToken)
    req.URL.RawQuery = query.Encode()
    utils.Log.Println("Sending request at : ", req.URL.String())
    resp, err := client.Do(req)
    defer resp.Body.Close()
    if err != nil {
       return nil, err
    }
    body, err := ioutil.ReadAll(resp.Body)
    if resp.StatusCode >= 400 {
        utils.Log.Println(string(body))
        return nil, errors.New(http.StatusText(resp.StatusCode))
    }
    if err != nil {
       return nil, err
    }
    return body, nil
}
