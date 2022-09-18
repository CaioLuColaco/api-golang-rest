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
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	r.HandleFunc("/api/personalidades/{id}", controllers.UmaPersonalidade)
	log.Fatal(http.ListenAndServe(":2004", r))
}