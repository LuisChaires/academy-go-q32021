package controllers

import (
	"html/template"
	"net/http"

	"deliverables/common/constants"
	"deliverables/usecases"

	"github.com/gorilla/mux"
)

// Home - shows an hellow greet
func Home(w http.ResponseWriter, r *http.Request) {
	showData(w, constants.PageData{Message: "Hello World", Status: "Success"})
}

//GetById - Function to get pokemon by ID
func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, _ := usecases.GetPokemonById(id)
	showData(w, data)
}

//GetAll - Function to get all stored pokemons
func GetAll(w http.ResponseWriter, r *http.Request) {
	data, _ := usecases.GetAllPokemons()
	showData(w, data)
}

func showData(w http.ResponseWriter, data constants.PageData) {
	tmpl := template.Must(template.ParseFiles(constants.LayoutIndex))
	tmpl.Execute(w, data)
}
