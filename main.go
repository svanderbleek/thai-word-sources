package main

import (
	"./forvo"
	"./sources"
	"./thai2english"
	"html/template"
	"net/http"
)

func main() {
	server := server()
	server.ListenAndServe()
}

func server() *http.Server {
	routes := routes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	return server
}

func routes() *http.ServeMux {
	routes := http.NewServeMux()
	routes.HandleFunc("/", homePage)
	routes.HandleFunc("/word", wordPage)
	routes.HandleFunc("/style.css", styleSheet)
	return routes
}

func styleSheet(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "style.css")
}

func homePage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "home.html")
}

func wordPage(response http.ResponseWriter, request *http.Request) {
	query := request.FormValue("query")
	source := sources.Bundle(thai2english.Search, forvo.Search)
	word := source(query)
	word.Text = query
	wordTemplate := template.Must(template.ParseFiles("word.html"))
	wordTemplate.Execute(response, word)
}
