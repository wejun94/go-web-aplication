package controllers

import (
	"html/template"
	"net/http"
	"web-application/models"
)

var temp = template.Must(template.ParseGlob("templates/*index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)

}
