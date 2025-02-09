package constants

import "deliverables/entities"

const (
	CvsFile            = "files/commas_file.csv"
	PokemonApiEndPoint = "pokemon/{id}"
)

//PageData - Struct to send data to the view
type PageData struct {
	Data    map[int]entities.Pokemon
	Message string
	Status  int
}
