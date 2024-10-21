package controllers

import (
	"Aplicacao_Web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templatesAlunos = template.Must(template.ParseGlob("templates/aluno/*.html"))

func IndexAlunos(w http.ResponseWriter, r *http.Request) {
	todasOsAlunos := models.BuscaAluno()
	templatesAlunos.ExecuteTemplate(w, "AlunosIndex", todasOsAlunos)
}

func AdicionarAlunos(w http.ResponseWriter, r *http.Request) {
	templatesAlunos.ExecuteTemplate(w, "AdicionarAlunos", nil)
}

func InserirAlunos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		email := r.FormValue("email")

		models.CriarAluno(nome, email)
	}
	http.Redirect(w, r, "/", 301)
}

func DeletarAlunos(w http.ResponseWriter, r *http.Request) {
	idAluno := r.URL.Query().Get("id")
	models.DeletarAluno(idAluno)
	http.Redirect(w, r, "/", 301)
}

func EditarAlunos(w http.ResponseWriter, r *http.Request) {
	idAluno := r.URL.Query().Get("id")
	aluno := models.EditarAluno(idAluno)
	templatesAlunos.ExecuteTemplate(w, "EditarAlunos", aluno)
}

func AtualizarAlunos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		email := r.FormValue("email")

		idConvertida, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para INT:", err)
		}

		models.AtualizarAluno(idConvertida, nome, email)
	}
	http.Redirect(w, r, "/", 301)
}
