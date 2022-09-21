package routes

import (
	"log"
	"net/http"

	"github.com/CaioLuColaco/api-golang-rest/controllers"
	"github.com/CaioLuColaco/api-golang-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.UmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarUmaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarUmaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":2004", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}