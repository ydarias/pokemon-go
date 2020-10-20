package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/repositories"
)

type PokemonTypesController struct {
	PokemonTypesRepository repositories.PokemonTypesRepository
}

func (p PokemonTypesController) Get(c *gin.Context) {
	pokemonTypes := p.PokemonTypesRepository.FindTypes()
	c.JSON(200, toStringTypes(pokemonTypes))
}
