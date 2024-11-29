package models

import (
	"api-go-rest/db"
)

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

func Create(personalidade Personalidade) {
	db := db.Connect()
	statement, err := db.Prepare("INSERT INTO personalidades (nome, historia) VALUES ($1,$2)")
	if err == nil {
		_, err := statement.Exec(personalidade.Nome, personalidade.Historia)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	defer db.Close()
}

func Read() []Personalidade {
	db := db.Connect()
	result, err := db.Query("SELECT * FROM personalidades")
	var personalidades []Personalidade
	if err == nil {
		for result.Next() {
			personalidade := Personalidade{}
			var id int
			var nome, historia string

			err = result.Scan(&id, &nome, &historia)
			if err == nil {
				personalidade.Id = id
				personalidade.Nome = nome
				personalidade.Historia = historia
				personalidades = append(personalidades, personalidade)
			} else {
				panic(err.Error())
			}
		}
	} else {
		panic(err.Error())
	}
	defer db.Close()
	return personalidades
}

func Update(personalidade Personalidade) {
	db := db.Connect()
	statement, err := db.Prepare("UPDATE personalidades SET nome=$2, historia=$3 WHERE id=$1 ")
	if err == nil {
		_, err := statement.Exec(personalidade.Id, personalidade.Nome, personalidade.Historia)
		if err != nil {
			panic(err.Error())
		}
	} else {
		panic(err.Error())
	}
	defer db.Close()
}

func Delete(id string) {
	db := db.Connect()
	statement, err := db.Prepare("DELETE FROM personalidades WHERE id=$1")
	if err == nil {
		statement.Exec(id)
	} else {
		panic(err.Error())
	}
	defer db.Close()
}
