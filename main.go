package main

import (
	"fmt"
	"github.com/NicolasKonishi/go-rest-api/database"
	"github.com/NicolasKonishi/go-rest-api/models"
	"github.com/NicolasKonishi/go-rest-api/routes"
)

func main() {
	models.Personas = []models.Personalidades{
		{Id: 1, Nome: "nome1", Historia: "historia1"},
		{Id: 2, Nome: "nome2", Historia: "historia2"},
	}

	database.InitDB()

	fmt.Println("iniciando server...http://localhost:8080/")
	routes.HandleRequest()
}
