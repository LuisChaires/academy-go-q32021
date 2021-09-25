package services

import (
	"bufio"
	"deliverables/common/constants"
	"encoding/json"
	"os"
	"strings"
)

var mapPokemon map[int]constants.Pokemon

func readCvsFile() (*os.File, error) {
	file, err := os.Open(constants.CvsFile)
	return file, err
}

//This function returns all the pokemons
func GetAllPokemons() (constants.PageData, error) {
	mapPokemon = make(map[int]constants.Pokemon)
	file, err := readCvsFile()
	if err != nil {
		file.Close()
		return constants.PageData{Message: "CSV File Error", Status: "409"}, err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		file.Close()
		return constants.PageData{Message: "File Read Error", Status: "409"}, err
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		mapPokemon[i] = constants.Pokemon{
			Id:       array[0],
			Name:     array[1],
			ImageUrl: array[2],
		}
	}

	jsonStr, err := json.MarshalIndent(mapPokemon, "", " ")
	if err != nil {
		file.Close()
		return constants.PageData{Message: "JSON parse error", Status: "409"}, err
	}

	file.Close()
	return constants.PageData{Message: string(jsonStr), Status: "200"}, err
}

//This function returns one pokemon by ID
func GetPokemonById(id string) (constants.PageData, error) {
	pokemon := constants.Pokemon{}
	file, err := readCvsFile()
	if err != nil {
		file.Close()
		return constants.PageData{Message: "CSV File Error", Status: "409"}, err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		file.Close()
		return constants.PageData{Message: "File Read Error", Status: "409"}, err
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		if array[0] == id {
			pokemon = constants.Pokemon{
				Id:       array[0],
				Name:     array[1],
				ImageUrl: array[2],
			}
			break
		}
	}

	if pokemon == (constants.Pokemon{}) {
		file.Close()
		return constants.PageData{Message: "No Data Found", Status: "404"}, err
	}

	jsonStr, err := json.MarshalIndent(pokemon, "", " ")
	if err != nil {
		file.Close()
		return constants.PageData{Message: "JSON parse error", Status: "409"}, err
	}

	file.Close()
	return constants.PageData{Message: string(jsonStr), Status: "200"}, err
}
