package controllers

import (
	"gin-api-rest/database"
	"gin-api-rest/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// ExibeTodosAlunos godoc
// @Summary alunos
// @Schemes
// @Description Exibe todos os alunos criados
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	if len(alunos) > 0 {
		c.JSON(http.StatusOK, alunos)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Empty": "Nenhum aluno foi cadastrado"})
	}
}

func ExibeAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID != 0 {
		c.JSON(http.StatusFound, aluno)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
		return
	}
}

func ExibeAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno) // armazena as informações encontradas no endereço de memória de aluno na linha 31
	if aluno.ID != 0 {
		c.JSON(http.StatusFound, aluno)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Not Found": "Aluno não encontrado"})
	}
}

func ApagaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"Aluno removido": id})
}

func AtualizaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err == nil {
		if err := models.Validate(&aluno); err == nil {
			database.DB.Model(&aluno).UpdateColumns(aluno)
			c.JSON(http.StatusOK, aluno)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Validation": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Saudações " + nome + "! Como vai?",
	})
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err == nil {
		if err := models.Validate(&aluno); err == nil {
			database.DB.Create(&aluno)
			c.JSON(http.StatusCreated, aluno)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"Validation": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
}

func IndexPage(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func NotFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
