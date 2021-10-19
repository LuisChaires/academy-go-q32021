package usecases

import (
	"errors"

	"deliverables/entities"
)

type services interface {
	GetAllPokemons() (map[int]entities.Pokemon, error)
	GetPokemonById(id string) (map[int]entities.Pokemon, error)
	GetPokemonFromAPI(id string) (entities.Pokemon, error)
	StorePokemon(pokemon entities.Pokemon) error
	GetConcurrently(pokemons map[int]entities.Pokemon, itemType string, items, ipw int) (map[int]entities.Pokemon, error)
}

//New - Function to get new usecases object
func New(s services) usecase {
	return usecase{s}
}

type usecase struct {
	service services
}

//GetAllPokemons - Function to get all stored pokemons
func (u usecase) GetAllPokemons() (map[int]entities.Pokemon, error) {
	return u.service.GetAllPokemons()
}

//GetPokemonById - Function to get pokemon by ID
func (u usecase) GetPokemonById(id string) (map[int]entities.Pokemon, error) {
	return u.service.GetPokemonById(id)
}

//GetPokemonFromAPI - Function to get data from external API
func (u usecase) GetPokemonFromAPI(id string) (map[int]entities.Pokemon, error) {
	pokemon, err := u.service.GetPokemonFromAPI(id)
	pokemonMap := make(map[int]entities.Pokemon)

	if err != nil {
		return nil, err
	}

	err = u.service.StorePokemon(pokemon)
	if err != nil {
		return nil, err
	}

	pokemonMap[0] = pokemon
	return pokemonMap, nil

}

//Concurrrency - Function to get data concurrently from CSV
func (u usecase) GetConcurrently(itemType string, items, ipw int) (map[int]entities.Pokemon, error) {
	var pokemons, err = u.service.GetAllPokemons()

	if err != nil {
		return nil, err
	}

	if len(pokemons) <= 0 {
		return nil, errors.New("no data found")
	}

	return u.service.GetConcurrently(pokemons, itemType, items, ipw)
}
