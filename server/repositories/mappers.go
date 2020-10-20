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
		Id:                    pokemonEntity.Identifier,
		Name:                  pokemonEntity.Name,
		Classification:        pokemonEntity.Classification,
		Types:                 toPokemonTypes(pokemonEntity.Types),
		Resistant:             toPokemonAttackTypes(pokemonEntity.Resistant),
		Weaknesses:            toPokemonAttackTypes(pokemonEntity.Weaknesses),
		MaxWeight:             pokemonEntity.MaxWeight,
		MinWeight:             pokemonEntity.MinWeight,
		MaxHeight:             pokemonEntity.MaxHeight,
		MinHeight:             pokemonEntity.MinHeight,
		FleeRate:              pokemonEntity.FleeRate,
		EvolutionRequirements: toPokemonEvolutionRequirements(pokemonEntity.EvolutionRequirements),
		Evolutions:            toEvolutionReferences(pokemonEntity.Evolutions),
		PreviousEvolutions:    toEvolutionReferences(pokemonEntity.PreviousEvolutions),
		MaxCP:                 pokemonEntity.MaxCP,
		MaxHP:                 pokemonEntity.MaxHP,
		FastAttacks:           toPokemonAttacks(pokemonEntity.FastAttacks),
		SpecialAttacks:        toPokemonAttacks(pokemonEntity.SpecialAttacks),
	}
}

func toPokemonTypes(entityTypes []db.PokemonEntityType) []PokemonType {
	var types []PokemonType

	for _, entityType := range entityTypes {
		types = append(types, PokemonType{entityType.Name})
	}

	return types
}

func toPokemonAttackTypes(entityAttacks []db.PokemonEntityAttackType) []PokemonAttackType {
	var attacks []PokemonAttackType

	for _, entityAttack := range entityAttacks {
		attacks = append(attacks, PokemonAttackType{entityAttack.Name})
	}

	return attacks
}

func toPokemonEvolutionRequirements(entityRequirements db.PokemonEntityEvolutionRequirements) EvolutionRequirements {
	return EvolutionRequirements{
		Name:   entityRequirements.Name,
		Amount: entityRequirements.Amount,
	}
}

func toEvolutionReferences(pokemons []*db.PokemonEntity) []EvolutionReference {
	var evolutionReferences []EvolutionReference

	for _, pokemon := range pokemons {
		evolutionReferences = append(evolutionReferences, EvolutionReference{
			Id:   pokemon.Identifier,
			Name: pokemon.Name,
		})
	}

	return evolutionReferences
}

func toPokemonAttacks(attacks []db.PokemonEntityAttack) []PokemonAttack {
	var pokemonAttacks []PokemonAttack

	for _, attack := range attacks {
		pokemonAttacks = append(pokemonAttacks, PokemonAttack{
			Name:   attack.Name,
			Type:   attack.Type.Name,
			Damage: attack.Damage,
		})
	}

	return pokemonAttacks
}
