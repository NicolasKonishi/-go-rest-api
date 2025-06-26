package main

import (
	"fmt"
	"github.com/NicolasKonishi/go-rest-api/routes"
)

func main() {
	fmt.Println("iniciando serveer...http://localhost:8080/")
	routes.HandleRequest()
}
