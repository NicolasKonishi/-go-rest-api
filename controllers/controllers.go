package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/NicolasKonishi/go-rest-api/database"
	"github.com/NicolasKonishi/go-rest-api/models"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Home Page")
}
func AllPersonas(w http.ResponseWriter, req *http.Request) {
	var p []models.Personalidades
	//DB.Find(&p) -> encontra no endereço de memória a var p
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
func GetPersona(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var p models.Personalidades
	//First(&p, id) -> retorna o primeiro que encontra de acordo com o id
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}
