package routes

import (
	"net/http"
	"web-application/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
