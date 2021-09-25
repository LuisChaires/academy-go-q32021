package usecases

import (
	"deliverables/common/constants"
	"deliverables/services"
)

func GetAllPokemons() (constants.PageData, error) {
	return services.GetAllPokemons()
}

func GetPokemonById(id string) (constants.PageData, error) {
	return services.GetPokemonById(id)
}
