package controllers

import (
	"Aplicacao_Web/models"
	"log"
	"net/http"
	"strconv"
)

func IndexMatriculas(w http.ResponseWriter, r *http.Request) {
	todasAsMatriculas := models.BuscaMatricula()
	templates.ExecuteTemplate(w, "MatriculasIndex", todasAsMatriculas)
}

func AdicionarMatriculas(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Adicionar", nil)
}

func InserirMatriculas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		quantidadeInt, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.CriarProduto(nome, descricao, precoFloat, quantidadeInt)
	}
	http.Redirect(w, r, "/", 301)
}

func DeletarMatriculas(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletarProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func EditarMatriculas(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditarProduto(idProduto)
	templates.ExecuteTemplate(w, "Editar", produto)
}

func AtualizarMatriculas(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertida, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do Id para INT:", err)
		}
		precoConvertida, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preço para Float64:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade para INT:", err)
		}

		models.AtualizarProduto(idConvertida, nome, descricao, precoConvertida, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}
