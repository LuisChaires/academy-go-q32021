package usecases

import (
	"deliverables/entities"
)

type services interface {
	GetAllPokemons() (map[int]entities.Pokemon, error)
	GetPokemonById(id string) (map[int]entities.Pokemon, error)
	GetPokemonFromAPI(id string) (entities.Pokemon, error)
	StorePokemon(pokemon entities.Pokemon) error
}

func New(s services) usecase {
	return usecase{s}
}

type usecase struct {
	service services
}

func (u usecase) GetAllPokemons() (map[int]entities.Pokemon, error) {
	return u.service.GetAllPokemons()
}

func (u usecase) GetPokemonById(id string) (map[int]entities.Pokemon, error) {
	return u.service.GetPokemonById(id)
}

func (u usecase) GetPokemonFromAPI(id string) (map[int]entities.Pokemon, error) {
	pokemon, err := u.service.GetPokemonFromAPI(id)
	pokemonMap := make(map[int]entities.Pokemon)

	if err != nil {
		return make(map[int]entities.Pokemon), err
	}

	err = u.service.StorePokemon(pokemon)
	if err != nil {
		return make(map[int]entities.Pokemon), err
	}

	pokemonMap[0] = pokemon
	return pokemonMap, nil

}
