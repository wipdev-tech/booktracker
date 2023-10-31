package main

import (
	"html/template"
	"net/http"
)

func main() {
	handleCSS := http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css/")))

	http.HandleFunc("/", handleIndex)
	http.Handle("/css/", handleCSS)
	http.ListenAndServe(":8000", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(
		template.ParseFiles("templates/index.html"),
	)

	tmpl.Execute(w, nil)
}
