package routes

import (
	"github.com/NicolasKonishi/go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
