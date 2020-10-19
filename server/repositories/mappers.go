package repositories

import "github.com/ydarias/pokemon-go/server/db"

func toPokemons(pokemonEntities []db.PokemonEntity) []Pokemon {
	var pokemons []Pokemon

	for _, pokemonEntity := range pokemonEntities {
		pokemons = append(pokemons, toPokemon(pokemonEntity))
	}

	return pokemons
}

func toPokemon(pokemonEntity db.PokemonEntity) Pokemon {
	return Pokemon{
		Id:             pokemonEntity.Identifier,
		Name:           pokemonEntity.Name,
		Classification: pokemonEntity.Classification,
		Types:          toPokemonTypes(pokemonEntity.Types),
		Resistant:      toPokemonAttacks(pokemonEntity.Resistant),
		Weaknesses:     toPokemonAttacks(pokemonEntity.Weaknesses),
		MaxWeight:      pokemonEntity.MaxWeight,
		MinWeight:      pokemonEntity.MinWeight,
		MaxHeight:      pokemonEntity.MaxHeight,
		MinHeight:      pokemonEntity.MinHeight,
		FleeRate:       pokemonEntity.FleeRate,
		MaxCP:          pokemonEntity.MaxCP,
		MaxHP:          pokemonEntity.MaxHP,
	}
}

func toPokemonTypes(entityTypes []db.PokemonEntityType) []PokemonType {
	var types []PokemonType

	for _, entityType := range entityTypes {
		types = append(types, PokemonType{entityType.Name})
	}

	return types
}

func toPokemonAttacks(entityAttacks []db.PokemonEntityAttack) []PokemonAttack {
	var attacks []PokemonAttack

	for _, entityAttack := range entityAttacks {
		attacks = append(attacks, PokemonAttack{entityAttack.Name})
	}

	return attacks
}
