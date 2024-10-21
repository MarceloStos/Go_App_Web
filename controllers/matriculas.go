package controllers

import (
	"Aplicacao_Web/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templatesMatriculas = template.Must(template.ParseGlob("templates/matricula/*.html"))

func IndexMatriculas(w http.ResponseWriter, r *http.Request) {
	todasAsMatriculas := models.BuscaMatricula()
	templatesMatriculas.ExecuteTemplate(w, "MatriculasIndex", todasAsMatriculas)
}

func AdicionarMatriculas(w http.ResponseWriter, r *http.Request) {
	templatesMatriculas.ExecuteTemplate(w, "AdicionarMatriculas", nil)
}

func InserirMatriculas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idAluno := r.FormValue("idAluno")
		idDisciplina := r.FormValue("idDisciplina")
		dataMatricula := r.FormValue("dataMatricula")

		idAlunoInt, err := strconv.Atoi(idAluno)

		if err != nil {
			log.Println("Erro na conversão do ID do aluno", err)
		}
		idDisciplinaInt, err := strconv.Atoi(idDisciplina)

		if err != nil {
			log.Println("Erro na conversão do ID da disciplina", err)
		}

		models.CriarMatricula(idAlunoInt, idDisciplinaInt, dataMatricula)
	}
	http.Redirect(w, r, "/matriculas", 301)
}

func DeletarMatriculas(w http.ResponseWriter, r *http.Request) {
	idAluno := r.URL.Query().Get("IdAluno")
	idDisciplina := r.URL.Query().Get("IdDisciplina")
	dataMatricula := r.URL.Query().Get("DataMatricula")
	models.DeletarMatricula(idAluno, idDisciplina, dataMatricula)
	http.Redirect(w, r, "/matriculas", 301)
}

func EditarMatriculas(w http.ResponseWriter, r *http.Request) {
	idAluno := r.URL.Query().Get("IdAluno")
	idDisciplina := r.URL.Query().Get("IdDisciplina")
	dataMatricula := r.URL.Query().Get("DataMatricula")

	matricula := models.EditarMatricula(idAluno, idDisciplina, dataMatricula)
	templatesMatriculas.ExecuteTemplate(w, "EditarMatriculas", matricula)
}

func AtualizarMatriculas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idAluno := r.FormValue("idAluno")
		idDisciplina := r.FormValue("idDisciplina")
		dataMatricula := r.FormValue("dataMatricula")

		idAlunoInt, err := strconv.Atoi(idAluno)

		if err != nil {
			log.Println("Erro na conversão do ID do aluno", err)
		}
		idDisciplinaInt, err := strconv.Atoi(idDisciplina)

		if err != nil {
			log.Println("Erro na conversão do ID da disciplina", err)
		}

		alunoIDAntigo := r.FormValue("idAlunoAntigo")           // Pegue do formulário
		disciplinaIDAntigo := r.FormValue("idDisciplinaAntigo") // Pegue do formulário

		alunoIDAntigoInt, err := strconv.Atoi(alunoIDAntigo)
		if err != nil {
			log.Println("Erro na conversão do ID antigo do aluno:", err)
		}

		disciplinaIDAntigoInt, err := strconv.Atoi(disciplinaIDAntigo)
		if err != nil {
			log.Println("Erro na conversão do ID antigo da disciplina:", err)
		}

		models.AtualizarMatricula(idAlunoInt, idDisciplinaInt, dataMatricula, alunoIDAntigoInt, disciplinaIDAntigoInt)
	}
	http.Redirect(w, r, "/matriculas", 301)
}
