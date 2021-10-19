package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"deliverables/common/constants"
	"deliverables/entities"

	"github.com/gorilla/mux"
)

type usecase interface {
	GetAllPokemons() (map[int]entities.Pokemon, error)
	GetPokemonById(id string) (map[int]entities.Pokemon, error)
	GetPokemonFromAPI(id string) (map[int]entities.Pokemon, error)
	GetConcurrently(itemType string, items, ipw int) (map[int]entities.Pokemon, error)
}

type controller struct {
	usecase usecase
}

//New - function to get new controller object
func New(u usecase) controller {
	return controller{u}
}

// Home - shows an hello greet
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

	if id == "" {
		showData(w, constants.PageData{Message: "Bad Request", Status: http.StatusBadRequest})
		return
	}

	data, err := c.usecase.GetPokemonFromAPI(id)
	if err != nil {
		showData(w, constants.PageData{Message: err.Error(), Status: http.StatusNotFound})
	} else {
		showData(w, constants.PageData{Data: data, Status: http.StatusOK})
	}
}

//Concurrrency - Function to get data concurrently from CSV
func (c controller) Concurrrency(w http.ResponseWriter, r *http.Request) {
	itemType := r.URL.Query().Get("type")
	items := r.URL.Query().Get("items")
	ipw := r.URL.Query().Get("items_per_worker")

	if itemType == "" || items == "" || ipw == "" {
		showData(w, constants.PageData{Message: "Missing Params", Status: http.StatusBadRequest})
		return
	}

	itemsQ, errQ := strconv.Atoi(items)
	itemsPW, errPW := strconv.Atoi(ipw)

	if errQ != nil || errPW != nil || (!strings.EqualFold(itemType, "even") && !strings.EqualFold(itemType, "odd")) {
		showData(w, constants.PageData{Message: "Invalid Params", Status: http.StatusBadRequest})
		return
	}

	data, err := c.usecase.GetConcurrently(itemType, itemsQ, itemsPW)
	if err != nil {
		showData(w, constants.PageData{Message: err.Error(), Status: http.StatusNotFound})
	} else {
		showData(w, constants.PageData{Data: data, Status: http.StatusOK})
	}
}

func showData(w http.ResponseWriter, data constants.PageData) {
	jsonStr, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(w, string(jsonStr))
	}
}
