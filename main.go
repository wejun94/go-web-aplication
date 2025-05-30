package main

import (
	"net/http"
	"web-application/routes"
)

func main() {
	//quando o programa iniciar, ele vai buscar da pasta routes o que deve ser importado
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
