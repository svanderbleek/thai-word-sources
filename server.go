package main

import (
	"./forvo"
	"./sources"
	"./thai2english"
	"./thailanguage"
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
	http.ServeFile(response, request, "view/style.css")
}

func homePage(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "view/home.html")
}

func wordPage(response http.ResponseWriter, request *http.Request) {
	query := request.FormValue("query")
	source := sources.Bundle(thai2english.Search, forvo.Search, thailanguage.Search)
	word := source(query)
	word.Text = query
	wordTemplate := template.Must(template.ParseFiles("view/word.html"))
	wordTemplate.Execute(response, word)
}
