package routes

import (
	"log"
	"net/http"

	"deliverables/controllers"

	"github.com/gorilla/mux"
)

//NewRouter - Function to set the routes
func NewRouter() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/all", controllers.GetAll)
	router.HandleFunc("/pokemon/{id}", controllers.GetById)

	log.Fatal(http.ListenAndServe(":10000", router))
}
