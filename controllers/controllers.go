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
	//DB.Find(&p) -> encontra no endereço de memória a var ps
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

func CreatePersona(w http.ResponseWriter, req *http.Request) {
	var p models.Personalidades
	//NewDecoder -> utiliza decoder ao inves de encoder para transformar json em linguagem que o go entende para salvar no banco
	json.NewDecoder(req.Body).Decode(&p)
	database.DB.Create(&p)
	// retorna a personalidade que foi criada
	json.NewEncoder(w).Encode(p)
}

func DeletePersona(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var p models.Personalidades

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditPersona(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var p models.Personalidades
	database.DB.First(&p, id)

	json.NewDecoder(req.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)

}
