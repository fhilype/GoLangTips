package main

import (
	"api_go/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/contacts", handlers.ListContacts)
	router.POST("/contacts", handlers.AddContact)
	router.PUT("/contacts/:id", handlers.EditContact)
	router.DELETE("/contacts/:id", handlers.RemoveContact)

	router.Run("localhost:8080")
}
