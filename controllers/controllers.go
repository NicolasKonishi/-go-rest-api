package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/NicolasKonishi/go-rest-api/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func AllPersonas(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(models.Personas)
}
func GetPersona(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	for _, persona := range models.Personas {
		//if strconv.Itoa(persona.Id) -> converte para string
		if strconv.Itoa(persona.Id) == id {
			json.NewEncoder(w).Encode(persona)
		}

	}
}
