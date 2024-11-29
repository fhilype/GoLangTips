package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	routes := mux.NewRouter()
	routes.Use(middleware.ContentTypeMiddleware)
	routes.HandleFunc("/", controllers.Home)
	routes.HandleFunc("/personalidades", controllers.Personalidades).Methods("GET")
	routes.HandleFunc("/personalidades", controllers.NovaPersonalidade).Methods("POST")
	routes.HandleFunc("/personalidade/{id}", controllers.Personalidade).Methods("GET")
	routes.HandleFunc("/personalidade/{id}", controllers.AtualizarPersonalidade).Methods("PUT")
	routes.HandleFunc("/personalidade/{id}", controllers.ApagarPersonalidade).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(routes)))
}
