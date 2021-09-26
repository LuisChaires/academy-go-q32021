package services

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"deliverables/common/constants"
)

var mapPokemon map[int]constants.Pokemon

func readCvsFile() (*os.File, error) {
	file, err := os.Open(constants.CvsFile)
	return file, err
}

//GetAllPokemons - This function returns all the pokemons
func GetAllPokemons() (constants.PageData, error) {
	mapPokemon = make(map[int]constants.Pokemon)
	file, err := readCvsFile()
	if err != nil {
		file.Close()
		return constants.PageData{Message: "CSV File Error", Status: constants.InternalError}, err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		file.Close()
		return constants.PageData{Message: "File Read Error", Status: constants.InternalError}, err
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		mapPokemon[i] = constants.Pokemon{
			ID:       array[0],
			Name:     array[1],
			ImageUrl: array[2],
		}
	}

	jsonStr, err := json.MarshalIndent(mapPokemon, "", " ")
	if err != nil {
		file.Close()
		return constants.PageData{Message: "JSON parse error", Status: constants.InternalError}, err
	}

	file.Close()
	return constants.PageData{Message: string(jsonStr), Status: constants.Success}, err
}

//GetPokemonById - This function returns one pokemon by ID
func GetPokemonById(id string) (constants.PageData, error) {
	pokemon := constants.Pokemon{}
	file, err := readCvsFile()
	if err != nil {
		file.Close()
		return constants.PageData{Message: "CSV File Error", Status: constants.InternalError}, err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		file.Close()
		return constants.PageData{Message: "File Read Error", Status: constants.InternalError}, err
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		if array[0] == id {
			pokemon = constants.Pokemon{
				ID:       array[0],
				Name:     array[1],
				ImageUrl: array[2],
			}
			break
		}
	}

	if pokemon == (constants.Pokemon{}) {
		file.Close()
		return constants.PageData{Message: "No Data Found", Status: constants.NotFound}, err
	}

	jsonStr, err := json.MarshalIndent(pokemon, "", " ")
	if err != nil {
		file.Close()
		return constants.PageData{Message: "JSON parse error", Status: constants.InternalError}, err
	}

	file.Close()
	return constants.PageData{Message: string(jsonStr), Status: constants.Success}, err
}
