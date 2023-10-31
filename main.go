package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8000", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, nil)
}
