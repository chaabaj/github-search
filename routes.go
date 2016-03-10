package main

import (
       "fmt"
       "net/http"
       "html/template"
       "github.com/chaabaj/github-search/utils"
       "github.com/chaabaj/github-search/service"
)

// Index route handler
// It load the homepage of the application
func baseRouteHandler(res http.ResponseWriter, req *http.Request) {
     utils.Log.Println("Load index.html")
     fmt.Fprintf(res, string(utils.ServeFile("templates/index.html")))
}

// Route for searching repositories it look
// if user try to send a request with search form value is empty
// It will redirect to the index page
func searchRepositoriesRouteHandler(res http.ResponseWriter, req *http.Request) {
    utils.Log.Println("Load data")
    if len(req.FormValue("search")) == 0 {
        http.Redirect(res, req, "/", 301)
    } else {
        repositories, err := service.SearchRepositories(req.FormValue("search"))
        if err != nil {
            utils.Log.Println("Cannot retreive data : " + err.Error())
            tpl, _ := template.ParseFiles("templates/error.html")
            tpl.Execute(res, err.Error())
        } else {
            tpl, _ := template.ParseFiles("templates/search.html")
            tpl.Execute(res, repositories)
        }
    }
}

// Register route handlers to the serve
// It register also a handler to handle static files like js, img, css,...
func RegisterRoutes() {
     http.HandleFunc("/", baseRouteHandler)
     utils.Log.Println("Register route handlers")
     http.HandleFunc("/search", searchRepositoriesRouteHandler)
     http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
         utils.Log.Println("Load file : " + r.URL.Path)
         http.ServeFile(w, r, r.URL.Path[1:])
    })
}
