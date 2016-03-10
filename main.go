package main

import (
       "net/http"
       "github.com/chaabaj/github-search/utils"
)

func main() {
     utils.Log.Println("Starting server on port 8080")
     RegisterRoutes()
     utils.Log.Println("Route handlers registered")
     http.ListenAndServe(":8080", nil)
}
