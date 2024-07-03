package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/value; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/value; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:rahulkottak@gmail.com\">rahulkottak</a>.")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	
}

type Router struct{}


func (router Router) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		http.Error(w, "not found", 404)
	}
}

func main() {
	var router Router
	fmt.Println("Started")
	http.ListenAndServe(":3000", router)
}
