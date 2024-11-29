package controllers

import (
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Personalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Read())
}

func Personalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Read() {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}

func NovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	models.Create(personalidade)
	personalidades := models.Read()
	personalidade = personalidades[len(personalidades)-1]
	json.NewEncoder(w).Encode(personalidade)
}

func ApagarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	models.Delete(id)
}

func AtualizarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)
	personalidade.Id, _ = strconv.Atoi(id)
	models.Update(personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
