package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/repositories"
)

func Pokemons(pokemonRepository repositories.PokemonRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		pokemons := pokemonRepository.FindPokemons()
		c.JSON(200, toPokemonsView(pokemons))
	}
}
