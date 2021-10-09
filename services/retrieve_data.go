package services

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"deliverables/common/constants"
	"deliverables/entities"

	cError "github.com/coreos/etcd/error"
	"gopkg.in/resty.v1"
)

type service struct {
	client *resty.Client
}

func New(host string, timeout time.Duration) (service, error) {
	client := resty.New().
		SetHostURL(host).
		SetTimeout(timeout).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			if r.IsSuccess() {
				return nil
			}

			return cError.NewError(r.StatusCode(), "error", 0)
		})

	return service{client}, nil

}

func readCvsFile() (*os.File, error) {
	file, err := os.Open(constants.CvsFile)
	return file, err
}

//GetAllPokemons - This function returns all the pokemons
func (s service) GetAllPokemons() (map[int]entities.Pokemon, error) {
	var mapPokemon = make(map[int]entities.Pokemon)
	file, err := readCvsFile()
	if err != nil {
		return mapPokemon, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanErr := scanner.Err(); scanErr != nil {
		return mapPokemon, scanErr
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		mapPokemon[i] = entities.Pokemon{
			ID:       array[0],
			Name:     array[1],
			ImageUrl: array[2],
		}
	}

	if err != nil {
		return mapPokemon, err
	}

	return mapPokemon, nil
}

//GetPokemonById - This function returns one pokemon by ID
func (s service) GetPokemonById(id string) (map[int]entities.Pokemon, error) {
	var mapPokemon = make(map[int]entities.Pokemon)

	file, err := readCvsFile()
	if err != nil {
		return mapPokemon, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanErr := scanner.Err(); scanErr != nil {
		return mapPokemon, scanErr
	}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		array := strings.Split(row, ",")
		if array[0] == id {
			mapPokemon[i] = entities.Pokemon{
				ID:       array[0],
				Name:     array[1],
				ImageUrl: array[2],
			}
			break
		}
	}

	if len(mapPokemon) == 0 {
		return mapPokemon, errors.New("no data")
	}

	if err != nil {
		file.Close()
		return mapPokemon, err
	}
	return mapPokemon, nil
}

func (s service) StorePokemon(pokemon entities.Pokemon) error {
	file, err := os.OpenFile(constants.CvsFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(
		pokemon.ID + "," +
			pokemon.Name + "," +
			pokemon.ImageUrl + "\n"); err != nil {
		return err
	}

	return nil
}

//getPokemonFromAPI - Send a request to external API
func (s service) GetPokemonFromAPI(id string) (entities.Pokemon, error) {
	out := &entities.Response{}
	var pokemon = entities.Pokemon{}

	resp, err := s.client.R().
		SetPathParams(map[string]string{"id": id}).
		SetHeader("Accept", "application/json").
		Get(constants.PokemonApi)

	if err != nil {
		return entities.Pokemon{}, err
	}

	body := resp.Body()
	if err := json.Unmarshal(body, out); err != nil {
		return entities.Pokemon{}, err
	}

	pokemon = entities.Pokemon{
		ID:       strconv.Itoa(out.ID),
		Name:     out.Name,
		ImageUrl: out.Sprites.Other.OfficialArtwork.ArtUrl,
	}

	return pokemon, nil
}
