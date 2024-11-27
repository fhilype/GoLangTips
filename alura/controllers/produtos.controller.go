package controllers

import (
	"alura/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(responseWriter http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(responseWriter, "index", models.BuscaTodosOsProdutos())
}

func New(responseWriter http.ResponseWriter, request *http.Request) {
	templates.ExecuteTemplate(responseWriter, "new", nil)
}

func Insert(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		preco, _err := strconv.ParseFloat(request.FormValue("preco"), 64)
		quantidade, err := strconv.Atoi(request.FormValue("quantidade"))
		if err != nil || _err != nil {
			log.Println("Erro ao converter:", err)
		}
		models.CriarNovoProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(responseWriter, request, "/", http.StatusMovedPermanently)
}

func Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	models.DeletaProduto(id)
	// http.Redirect(responseWriter, request, "/", 301) // literal
	http.Redirect(responseWriter, request, "/", http.StatusMovedPermanently)
}

func Edit(responseWriter http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	produto := models.EditaProduto(id)
	templates.ExecuteTemplate(responseWriter, "edit", produto)
}

func Update(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		id, err := strconv.Atoi(request.FormValue("id"))
		nome := request.FormValue("nome")
		descricao := request.FormValue("descricao")
		preco, _err := strconv.ParseFloat(request.FormValue("preco"), 64)
		quantidade, err_ := strconv.Atoi(request.FormValue("quantidade"))
		if err != nil || _err != nil || err_ != nil {
			log.Println("Erro ao converter:", err)
		}
		models.AtualizaProduto(id, nome, descricao, preco, quantidade)
	}
	http.Redirect(responseWriter, request, "/", http.StatusMovedPermanently)
}
