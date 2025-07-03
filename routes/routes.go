package routes

import (
	"github.com/NicolasKonishi/go-rest-api/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.AllPersonas)
	r.HandleFunc("/api/personas/{id}", controllers.GetPersona)

	log.Fatal(http.ListenAndServe(":8080", r))
}
