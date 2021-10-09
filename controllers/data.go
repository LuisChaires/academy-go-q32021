package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"deliverables/common/constants"
	"deliverables/entities"

	"github.com/gorilla/mux"
)

type usecase interface {
	GetAllPokemons() (map[int]entities.Pokemon, error)
	GetPokemonById(id string) (map[int]entities.Pokemon, error)
	GetPokemonFromAPI(id string) (map[int]entities.Pokemon, error)
}

type controller struct {
	usecase usecase
}

func New(u usecase) controller {
	return controller{u}
}

// Home - shows an hellow greet
func (c controller) Home(w http.ResponseWriter, r *http.Request) {
	showData(w, constants.PageData{Message: "Hello World", Status: http.StatusOK})
}

//GetById - Function to get pokemon by ID
func (c controller) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := c.usecase.GetPokemonById(id)

	if err != nil {
		showData(w, constants.PageData{Message: err.Error(), Status: http.StatusInternalServerError})
	} else {
		showData(w, constants.PageData{Data: data, Status: http.StatusOK})
	}
}

//GetAll - Function to get all stored pokemons
func (c controller) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := c.usecase.GetAllPokemons()
	if err != nil {
		showData(w, constants.PageData{Message: err.Error(), Status: http.StatusInternalServerError})
	} else {
		showData(w, constants.PageData{Data: data, Status: http.StatusOK})
	}
}

//GetAll - Function to get data from external API
func (c controller) GetFromAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	data, err := c.usecase.GetPokemonFromAPI(id)
	if err != nil {
		showData(w, constants.PageData{Message: err.Error(), Status: http.StatusInternalServerError})
	} else {
		showData(w, constants.PageData{Data: data, Status: http.StatusOK})
	}
}

func showData(w http.ResponseWriter, data constants.PageData) {
	jsonStr, _ := json.MarshalIndent(data, "", " ")
	fmt.Fprintf(w, string(jsonStr))
}
