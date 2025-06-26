package main

import (
	"fmt"
	"github.com/NicolasKonishi/go-rest-api/routes"
)

func main() {
	fmt.Println("iniciando server...http://localhost:8080/")
	routes.HandleRequest()
}
