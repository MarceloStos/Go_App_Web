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
	http.HandleFunc("/matriculas/adicionar", controllers.AdicionarMatriculas)
	http.HandleFunc("/matriculas/inserir", controllers.InserirMatriculas)
	http.HandleFunc("/matriculas/editar", controllers.EditarMatriculas)
	http.HandleFunc("/matriculas/atualizar", controllers.AtualizarMatriculas)
	http.HandleFunc("/matriculas/deletar", controllers.DeletarMatriculas)

	http.HandleFunc("/alunos", controllers.IndexAlunos)
	http.HandleFunc("/alunos/adicionar", controllers.AdicionarAlunos)
	http.HandleFunc("/alunos/inserir", controllers.InserirAlunos)
	http.HandleFunc("/alunos/editar", controllers.EditarAlunos)
	http.HandleFunc("/alunos/atualizar", controllers.AtualizarAlunos)
	http.HandleFunc("/alunos/deletar", controllers.DeletarAlunos)

	http.HandleFunc("/disciplinas", controllers.IndexDisciplinas)
	http.HandleFunc("/disciplinas/adicionar", controllers.AdicionarDisciplinas)
	http.HandleFunc("/disciplinas/inserir", controllers.InserirDisciplinas)
	http.HandleFunc("/disciplinas/editar", controllers.EditarDisciplinas)
	http.HandleFunc("/disciplinas/atualizar", controllers.AtualizarDisciplinas)
	http.HandleFunc("/disciplinas/deletar", controllers.DeletarDisciplinas)

}
