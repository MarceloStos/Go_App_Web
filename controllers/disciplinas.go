package controllers

import (
	"Aplicacao_Web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templatesDisciplinas = template.Must(template.ParseGlob("templates/disciplina/*.html"))

func IndexDisciplinas(w http.ResponseWriter, r *http.Request) {
	todasAsDisciplinas := models.BuscaDisciplina()
	templatesDisciplinas.ExecuteTemplate(w, "DisciplinasIndex", todasAsDisciplinas)
}

func AdicionarDisciplinas(w http.ResponseWriter, r *http.Request) {
	templatesDisciplinas.ExecuteTemplate(w, "AdicionarDisciplinas", nil)
}

func InserirDisciplinas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		models.CriarDisciplina(nome, descricao)
	}
	http.Redirect(w, r, "/", 301)
}

func DeletarDisciplinas(w http.ResponseWriter, r *http.Request) {
	idDisciplina := r.URL.Query().Get("id")
	models.DeletarDisciplina(idDisciplina)
	http.Redirect(w, r, "/", 301)
}

func EditarDisciplinas(w http.ResponseWriter, r *http.Request) {
	idDisciplina := r.URL.Query().Get("id")
	disciplina := models.EditarDisciplina(idDisciplina)
	templatesDisciplinas.ExecuteTemplate(w, "EditarDisciplinas", disciplina)
}

func AtualizarDisciplinas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")

		idConvertida, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para INT:", err)
		}

		models.AtualizarDisciplina(idConvertida, nome, descricao)
	}
	http.Redirect(w, r, "/", 301)
}
