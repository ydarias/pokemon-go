package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/repositories"
)

func GetPokemonTypes(pokemonTypesRepository repositories.PokemonTypesRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		pokemonTypes := pokemonTypesRepository.FindTypes()
		c.JSON(200, toStringArray(pokemonTypes))
	}
}

func toStringArray(types []repositories.PokemonType) []string {
	var result []string

	for _, aType := range types {
		result = append(result, aType.Name)
	}

	return result
}
