package main

import (
       "fmt"
       "net/http"
       "html/template"
       "github.com/chaabaj/github-search/utils"
       "github.com/chaabaj/github-search/service"
)

func baseRouteHandler(res http.ResponseWriter, req *http.Request) {
     utils.Log.Println("Load index.html")
     fmt.Fprintf(res, string(utils.ServeFile("templates/index.html")))
}

func searchRouteHandler(res http.ResponseWriter, req *http.Request) {
     utils.Log.Println("Load data")
     repositories, err := service.SearchRepositories(req.FormValue("search"))
     if err != nil {
         utils.Log.Println("Cannot retreive data : " + err.Error())
         fmt.Fprintf(res, err.Error())
     } else {
         tpl, err := template.ParseFiles("templates/search.html")
         if err != nil {
             http.Error(res, "Page not found", 404)
         } else {
             tpl.Execute(res, repositories)
         }
     }
}

func RegisterRoutes() {
     utils.Log.Println("Register route handlers")
     http.HandleFunc("/", baseRouteHandler)
     http.HandleFunc("/search", searchRouteHandler)
     http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
         utils.Log.Println("Load file : " + r.URL.Path)
         http.ServeFile(w, r, r.URL.Path[1:])
    })
}
