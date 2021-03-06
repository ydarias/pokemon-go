package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/repositories"
)

type PokemonsController struct {
	PokemonRepository repositories.PokemonRepository
}

func (p PokemonsController) Get(c *gin.Context) {
	pokemons := p.PokemonRepository.FindPokemons()
	c.JSON(200, toPokemonsView(pokemons))
}

func (p PokemonsController) GetById(c *gin.Context) {
	id := c.Param("pokemonId")
	pokemon := p.PokemonRepository.FindPokemonById(id)
	c.JSON(200, toPokemonView(pokemon))
}

func (p PokemonsController) GetByName(c *gin.Context) {
	name := c.Param("pokemonName")
	pokemon := p.PokemonRepository.FindPokemonByName(name)
	c.JSON(200, toPokemonView(pokemon))
}
