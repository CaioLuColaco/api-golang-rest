package routes

import (
	"log"
	"net/http"

	"github.com/CaioLuColaco/api-golang-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.UmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarUmaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditaPersonalidade).Methods("Put")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarUmaPersonalidade).Methods("Delete")
	log.Fatal(http.ListenAndServe(":2004", r))
}