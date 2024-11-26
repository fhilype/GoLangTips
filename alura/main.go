package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err == nil {
		return db
	} else {
		panic(err.Error())
	}
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func index(responseWriter http.ResponseWriter, request *http.Request) {
	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("SELECT * FROM produtos")
	var produtos []Produto
	if err == nil {
		for selectDeTodosOsProdutos.Next() {
			produto := Produto{}
			var id, quantidade int
			var nome, descricao string
			var preco float64

			err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err == nil {
				produto.Nome = nome
				produto.Descricao = descricao
				produto.Preco = preco
				produto.Quantidade = quantidade
				produtos = append(produtos, produto)
			} else {
				panic(err.Error())
			}
		}
	} else {
		panic(err.Error())
	}
	templates.ExecuteTemplate(responseWriter, "index", produtos)
	defer db.Close()
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}
