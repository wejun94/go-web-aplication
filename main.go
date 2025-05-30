package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*index.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)

}
