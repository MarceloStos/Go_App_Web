package controllers

import (
	"Aplicacao_Web/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func Adicionar(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "Adicionar", nil)
}
