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
	r.HandleFunc("/api/personas", controllers.AllPersonas).Methods("GET")
	r.HandleFunc("/api/personas/{id}", controllers.GetPersona).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
