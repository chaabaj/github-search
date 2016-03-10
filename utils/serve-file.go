package utils

import "io/ioutil"

func ServeFile(path string) []byte {
     body, err := ioutil.ReadFile(path)
     if err != nil {
         errBody, _ := ioutil.ReadFile("templates/404.html")
         return errBody
     }
     return body
}
