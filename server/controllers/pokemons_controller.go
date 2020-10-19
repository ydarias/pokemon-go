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

func toPokemonsView(pokemons []repositories.Pokemon) []PokemonView {
	var pokemonViews []PokemonView

	for _, pokemon := range pokemons {
		pokemonViews = append(pokemonViews, toPokemonView(pokemon))
	}

	return pokemonViews
}

func toPokemonView(pokemon repositories.Pokemon) PokemonView {
	return PokemonView{
		Id:             pokemon.Id,
		Name:           pokemon.Name,
		Classification: pokemon.Classification,
		Types:          toPokemonViewTypes(pokemon.Types),
	}
}

func toPokemonViewTypes(pokemonTypes []repositories.PokemonType) []string {
	var pokemonViewTypes []string

	for _, pokemonType := range pokemonTypes {
		pokemonViewTypes = append(pokemonViewTypes, pokemonType.Name)
	}

	return pokemonViewTypes
}
