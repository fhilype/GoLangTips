package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCPF)
	r.POST("/alunos", controllers.CriaAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.ApagaAlunoPorId)
	r.Run() // se não declarado, utiliza 8080 // r.Run(":5000")
}
