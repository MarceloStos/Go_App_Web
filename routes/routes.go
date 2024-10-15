package routes

import (
	"Aplicacao_Web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/adicionar", controllers.Adicionar)
	http.HandleFunc("/inserir", controllers.Inserir)
	http.HandleFunc("/deletar", controllers.Deletar)
	http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/update", controllers.Atualizar)

}
