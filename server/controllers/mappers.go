package controllers

import "github.com/ydarias/pokemon-go/server/repositories"

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
		Resistant:      toPokemonViewAttacks(pokemon.Resistant),
		Weaknesses:     toPokemonViewAttacks(pokemon.Weaknesses),
	}
}

func toPokemonViewTypes(pokemonTypes []repositories.PokemonType) []string {
	var pokemonViewTypes []string

	for _, pokemonType := range pokemonTypes {
		pokemonViewTypes = append(pokemonViewTypes, pokemonType.Name)
	}

	return pokemonViewTypes
}

func toPokemonViewAttacks(pokemonAttacks []repositories.PokemonAttack) []string {
	var pokemonViewAttacks []string

	for _, pokemonAttack := range pokemonAttacks {
		pokemonViewAttacks = append(pokemonViewAttacks, pokemonAttack.Name)
	}

	return pokemonViewAttacks
}