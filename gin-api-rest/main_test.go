package main

import (
	"bytes"
	"encoding/json"
	"gin-api-rest/controllers"
	"gin-api-rest/database"
	"gin-api-rest/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func Routes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode) // declarado ? mostra os testes detalhados : mostra os testes resumidos
	r := gin.Default()
	return r
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "Aluno Teste",
		CPF:  "99999999999",
		RG:   "9999999",
	}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestStatusCodeSaudacao(t *testing.T) {
	r := Routes()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/Fhilype", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	/*
		if res.Code != http.StatusOK {
			t.Fatalf("[Status error] Recebido: %d, Esperado: %d", res.Code, http.StatusOK)
		}
	*/
	assert.Equal(t, http.StatusOK, res.Code, "SHOULD EQUAL")
	mockRes := `{"API diz:":"Saudações Fhilype! Como vai?"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody), "SHOULD EQUAL")
}

func TestExibeTodosAlunos(t *testing.T) {
	database.Connect()
	CriaAlunoMock()
	r := Routes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	defer DeletaAlunoMock()
}

func TestExibeAlunoPorCPF(t *testing.T) {
	database.Connect()
	CriaAlunoMock()
	r := Routes()
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/99999999999", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusFound, res.Code)
	defer DeletaAlunoMock()
}

func TestExibeAlunoPorId(t *testing.T) {
	database.Connect()
	CriaAlunoMock()
	r := Routes()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoMock models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoMock)
	assert.Equal(t, http.StatusFound, res.Code)
	assert.Equal(t, "Aluno Teste", alunoMock.Nome)
	assert.Equal(t, "99999999999", alunoMock.CPF)
	assert.Equal(t, "9999999", alunoMock.RG)
	defer DeletaAlunoMock()
}

func TestApagaAlunoPorId(t *testing.T) {
	database.Connect()
	CriaAlunoMock()
	r := Routes()
	r.DELETE("/alunos/:id", controllers.ApagaAlunoPorId)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAtualizaAluno(t *testing.T) {
	database.Connect()
	CriaAlunoMock()
	r := Routes()
	r.PATCH("/alunos/:id", controllers.AtualizaAluno)
	aluno := models.Aluno{
		Nome: "Aluno Teste",
		CPF:  "00000000000",
		RG:   "0000000",
	}
	alunoJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(alunoJson))
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	var alunoAtualizado models.Aluno
	json.Unmarshal(res.Body.Bytes(), &alunoAtualizado)
	assert.Equal(t, aluno.Nome, alunoAtualizado.Nome)
	assert.Equal(t, aluno.CPF, alunoAtualizado.CPF)
	assert.Equal(t, aluno.RG, alunoAtualizado.RG)
	defer DeletaAlunoMock()
}
