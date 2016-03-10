package main

import (
        "fmt"
        "os"
        "net/http"
        "github.com/chaabaj/github-search/utils"
)

func main() {
    if len(os.Getenv("GITHUB_AUTH_TOKEN")) == 0 {
        fmt.Println("Please set the variable GITHUB_AUTH_TOKEN and restart")
    } else {
        utils.Log.Println("Starting server on port 8080")
        RegisterRoutes()
        utils.Log.Println("Route handlers registered")
        http.ListenAndServe(":8080", nil)
    }
}
