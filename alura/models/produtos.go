package models

import (
	"alura/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
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
				produto.Id = id
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
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insertValues, err := db.Prepare("INSERT INTO produtos(nome, descricao, preco, quantidade) VALUES($1,$2,$3,$4)")
	if err == nil {
		insertValues.Exec(nome, descricao, preco, quantidade)
	} else {
		panic(err.Error())
	}
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()
	deleteById, err := db.Prepare("DELETE FROM produtos WHERE id=$1 ORDER BY id ASC")
	if err == nil {
		deleteById.Exec(id)
	} else {
		panic(err.Error())
	}
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("SELECT * FROM produtos WHERE id=$1", id)
	produto := Produto{}
	if err == nil {
		for produtoDoBanco.Next() {
			var id, quantidade int
			var nome, descricao string
			var preco float64

			err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
			if err == nil {
				produto.Id = id
				produto.Nome = nome
				produto.Descricao = descricao
				produto.Preco = preco
				produto.Quantidade = quantidade
			} else {
				panic(err.Error())
			}
		}
	} else {
		panic(err.Error())
	}
	defer db.Close()
	return produto
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	updateValues, err := db.Prepare("UPDATE produtos SET nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
	if err == nil {
		updateValues.Exec(nome, descricao, preco, quantidade, id)
	} else {
		panic(err.Error())
	}
	defer db.Close()
}
