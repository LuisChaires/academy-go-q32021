package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//En revision
func TestPokemonId(t *testing.T) {
	assert := assert.New(t)

	var tests = []struct {
		input string
		err   error
	}{
		{"1", nil},
	}

	for _, test := range tests {
		_, err := services.GetPokemonById(test.input)
		assert.Equal(err, test.err)
	}
}
