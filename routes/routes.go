package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type controller interface {
	Home(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetFromAPI(w http.ResponseWriter, r *http.Request)
	Concurrrency(w http.ResponseWriter, r *http.Request)
}

//New - Function to set the routes
func New(c controller) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", c.Home)
	router.HandleFunc("/all", c.GetAll)
	router.HandleFunc("/pokemon/{id}", c.GetById)
	router.HandleFunc("/pokemon/api/{id}", c.GetFromAPI)

	router.HandleFunc("/concurrency", c.Concurrrency)

	log.Fatal(http.ListenAndServe(":10000", router))
}
