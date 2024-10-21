package routes

import (
	"Aplicacao_Web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	//http.HandleFunc("/adicionar", controllers.Adicionar)
	//http.HandleFunc("/inserir", controllers.Inserir)
	//http.HandleFunc("/deletar", controllers.Deletar)
	//http.HandleFunc("/editar", controllers.Editar)
	http.HandleFunc("/update", controllers.Atualizar)
	http.HandleFunc("/matriculas", controllers.IndexMatriculas)

	http.HandleFunc("/disciplinas", controllers.IndexDisciplinas)
	http.HandleFunc("/disciplinas/adicionar", controllers.AdicionarDisciplinas)
	http.HandleFunc("/disciplinas/inserir", controllers.InserirDisciplinas)
	http.HandleFunc("/disciplinas/editar", controllers.EditarDisciplinas)
	http.HandleFunc("/disciplinas/atualizar", controllers.AtualizarDisciplinas)
	http.HandleFunc("/disciplinas/deletar", controllers.DeletarDisciplinas)

}
