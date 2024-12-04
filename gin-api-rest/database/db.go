package database

import (
	"gin-api-rest/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	connectionParams := "host=localhost user=postgres password=postgres dbname=alunos port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionParams))
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	DB.AutoMigrate(&models.Aluno{})
}
