package routes

import (
	"Aplicacao_Web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/adicionar", controllers.Adicionar)

}
