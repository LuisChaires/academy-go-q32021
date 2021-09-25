package controllers

import (
	"html/template"
	"net/http"
	"os"

	"deliverables/common/constants"
	"deliverables/usecases"

	"github.com/gorilla/mux"
)

// Home - Description
func Home(w http.ResponseWriter, r *http.Request) {
	showData(w, constants.PageData{Message: "Hello World", Status: "Success"})
}

//mover a useCase y Service
func GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, _ := usecases.GetPokemonById(id)
	showData(w, data)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	data, _ := usecases.GetAllPokemons()
	showData(w, data)
}

func ReadCvsFile() (*os.File, error) {
	file, err := os.Open(constants.CvsFile)
	return file, err
}

func showData(w http.ResponseWriter, data constants.PageData) {
	tmpl := template.Must(template.ParseFiles(constants.LayoutIndex))
	tmpl.Execute(w, data)
}
