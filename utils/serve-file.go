package utils

import "io/ioutil"

// Load file if it exist
// If it don't exist it load a 404.html page to display it to the user 
func ServeFile(path string) []byte {
     body, err := ioutil.ReadFile(path)
     if err != nil {
         errBody, _ := ioutil.ReadFile("templates/404.html")
         return errBody
     }
     return body
}
