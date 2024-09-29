package main

import (
	"Aplicacao_Web/routes"
	"net/http"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
