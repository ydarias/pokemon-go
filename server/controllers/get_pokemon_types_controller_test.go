package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/ydarias/pokemon-go/server/repositories"
	"net/http/httptest"
	"testing"
)

func TestGetPokemonTypes(t *testing.T) {
	pokemonTypeRepository := PokemonTypesMockRepository{}
	getPokemonTypesController := GetPokemonTypes(pokemonTypeRepository)
	t.Run("should return the types from the repository", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		getPokemonTypesController(c)

		assert.Equal(t, "[\"Foo\",\"Bar\"]", w.Body.String())
	})
}

type PokemonTypesMockRepository struct {
}

func (p PokemonTypesMockRepository) FindTypes() []repositories.PokemonType {
	return []repositories.PokemonType{
		{"Foo"},
		{"Bar"},
	}
}
