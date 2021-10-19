package usecases

import (
	"errors"
	"testing"

	"deliverables/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var pokemons = map[int]entities.Pokemon{
	0: {
		ID:   "1",
		Name: "bulbasaur",
	},
	1: {
		ID:   "4",
		Name: "charmander",
	},
	2: {
		ID:   "7",
		Name: "squirtle",
	},
}

var pokemon = map[int]entities.Pokemon{
	0: {
		ID:   "1",
		Name: "bulbasaur",
	},
}

type PokemonMock struct {
	mock.Mock
}

func (m PokemonMock) GetPokemonById(id string) (map[int]entities.Pokemon, error) {
	args := m.Called(id)
	return args.Get(0).(map[int]entities.Pokemon), args.Error(1)
}

func (m PokemonMock) GetAllPokemons() (map[int]entities.Pokemon, error) {
	args := m.Called()
	return args.Get(0).(map[int]entities.Pokemon), args.Error(1)
}

func (m PokemonMock) StorePokemon(pokemon entities.Pokemon) error {
	args := m.Called()
	return args.Error(1)
}

func (m PokemonMock) GetPokemonFromAPI(id string) (entities.Pokemon, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Pokemon), args.Error(1)
}

func (m PokemonMock) GetConcurrently(pokemons map[int]entities.Pokemon, itemType string, items, ipw int) (map[int]entities.Pokemon, error) {
	args := m.Called()
	return args.Get(0).(map[int]entities.Pokemon), args.Error(1)
}

func TestPokemonId(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       map[int]entities.Pokemon
		hasError       bool
		err            error
		ID             string
	}{
		{
			"Find result",
			1,
			pokemon,
			false,
			nil,
			"1",
		},
		{
			"No Data Found",
			0,
			map[int]entities.Pokemon{},
			true,
			errors.New("no data found"),
			"2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := PokemonMock{}
			mock.On("GetPokemonById", tc.ID).Return(tc.response, tc.err)
			service := New(mock)

			result, err := service.GetPokemonById(tc.ID)
			assert.EqualValues(t, tc.expectedLength, len(result))

			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}

		})
	}

}

func TestPokemonAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedLength int
		response       map[int]entities.Pokemon
		hasError       bool
		err            error
	}{
		{
			"Find result",
			3,
			pokemons,
			false,
			nil,
		},
		{
			"No Data Found",
			0,
			map[int]entities.Pokemon{},
			true,
			errors.New("no data found"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := PokemonMock{}
			mock.On("GetAllPokemons").Return(tc.response, tc.err)
			service := New(mock)

			result, err := service.GetAllPokemons()
			assert.EqualValues(t, tc.expectedLength, len(result))

			if tc.hasError {
				assert.EqualError(t, err, tc.err.Error())
			} else {
				assert.Nil(t, err)
			}

		})
	}

}
