package routes

import (
	"github.com/NicolasKonishi/go-rest-api/controllers"
	"github.com/NicolasKonishi/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personas", controllers.AllPersonas).Methods("GET")
	r.HandleFunc("/api/personas/{id}", controllers.GetPersona).Methods("GET")
	r.HandleFunc("/api/personas", controllers.CreatePersona).Methods("POST")
	r.HandleFunc("/api/personas/{id}", controllers.DeletePersona).Methods("DELETE")
	r.HandleFunc("/api/personas/{id}", controllers.EditPersona).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
