package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/ydarias/pokemon-go/server/repositories"
	"net/http/httptest"
	"testing"
)

func TestPokemonsController(t *testing.T) {

	pokemonRepository := PokemonMockRepository{}
	pokemonsController := PokemonsController{PokemonRepository: pokemonRepository}

	t.Run("should return a pokemon by Id", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = append(c.Params, gin.Param{
			Key:   "pokemonId",
			Value: "001",
		})
		pokemonsController.GetById(c)

		var result PokemonView
		json.Unmarshal(w.Body.Bytes(), &result)

		assert.Equal(t, "001", result.Id)
	})

	t.Run("should return a pokemon by name", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Params = append(c.Params, gin.Param{
			Key:   "pokemonName",
			Value: "bunny",
		})
		pokemonsController.GetByName(c)

		var result PokemonView
		json.Unmarshal(w.Body.Bytes(), &result)

		assert.Equal(t, "bunny", result.Name)
	})
}

type PokemonMockRepository struct{}

func (p PokemonMockRepository) FindPokemons() []repositories.Pokemon {
	return []repositories.Pokemon{}
}

func (p PokemonMockRepository) FindPokemonById(id string) repositories.Pokemon {
	return repositories.Pokemon{
		Id:             id,
		Name:           "Foo",
		Classification: "A classification",
	}
}

func (p PokemonMockRepository) FindPokemonByName(name string) repositories.Pokemon {
	return repositories.Pokemon{
		Id:             "001",
		Name:           name,
		Classification: "A classification",
	}
}
