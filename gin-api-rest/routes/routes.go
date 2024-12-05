package routes

import (
	"gin-api-rest/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func HandleRequests() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.LoadHTMLGlob("templates/*")   // indica o diretório das páginas html para serem carregadas
	r.Static("/assets", "./assets") // indica o diretório de estilos css para serem carregados
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCPF)
	r.POST("/alunos", controllers.CriaAluno)
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	r.DELETE("/alunos/:id", controllers.ApagaAlunoPorId)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", controllers.IndexPage)
	r.NoRoute(controllers.NotFoundPage) // quando nenhuma rota for encontrada
	r.Run()                             // se não declarado, utiliza 8080 // r.Run(":5000")
}
