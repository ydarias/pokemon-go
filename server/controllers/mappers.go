package controllers

import (
	"github.com/ydarias/pokemon-go/server/repositories"
	"strconv"
)

func toPokemonsView(pokemons []repositories.Pokemon) []PokemonView {
	var pokemonViews []PokemonView

	for _, pokemon := range pokemons {
		pokemonViews = append(pokemonViews, toPokemonView(pokemon))
	}

	return pokemonViews
}

func toPokemonView(pokemon repositories.Pokemon) PokemonView {
	weight := RangeView{
		Maximum: pokemon.MaxWeight,
		Minimum: pokemon.MinWeight,
	}
	height := RangeView{
		Maximum: pokemon.MaxHeight,
		Minimum: pokemon.MinHeight,
	}
	evolutionRequirements := EvolutionRequirements{
		Amount: pokemon.EvolutionRequirements.Amount,
		Name:   pokemon.EvolutionRequirements.Name,
	}
	return PokemonView{
		Id:                    pokemon.Id,
		Name:                  pokemon.Name,
		Classification:        pokemon.Classification,
		Types:                 toPokemonViewTypes(pokemon.Types),
		Resistant:             toPokemonViewAttacks(pokemon.Resistant),
		Weaknesses:            toPokemonViewAttacks(pokemon.Weaknesses),
		Weight:                weight,
		Height:                height,
		FleeRate:              pokemon.FleeRate,
		EvolutionRequirements: evolutionRequirements,
		MaxCP:                 pokemon.MaxCP,
		MaxHP:                 pokemon.MaxHP,
		Evolutions:            toPokemonEvolutions(pokemon.Evolutions),
		PreviousEvolutions:    toPokemonEvolutions(pokemon.PreviousEvolutions),
		Attacks:               toAttacks(pokemon),
	}
}

func toPokemonViewTypes(pokemonTypes []repositories.PokemonType) []string {
	var pokemonViewTypes []string

	for _, pokemonType := range pokemonTypes {
		pokemonViewTypes = append(pokemonViewTypes, pokemonType.Name)
	}

	return pokemonViewTypes
}

func toPokemonViewAttacks(pokemonAttacks []repositories.PokemonAttackType) []string {
	var pokemonViewAttacks []string

	for _, pokemonAttack := range pokemonAttacks {
		pokemonViewAttacks = append(pokemonViewAttacks, pokemonAttack.Name)
	}

	return pokemonViewAttacks
}

func toPokemonEvolutions(evolutions []repositories.EvolutionReference) []EvolutionReference {
	if len(evolutions) == 0 {
		return []EvolutionReference{}
	}

	var pokemonEvolutions []EvolutionReference

	for _, evolution := range evolutions {
		evolutionId, _ := strconv.Atoi(evolution.Id)
		pokemonEvolutions = append(pokemonEvolutions, EvolutionReference{
			Id:   evolutionId,
			Name: evolution.Name,
		})
	}

	return pokemonEvolutions
}

func toAttacks(pokemon repositories.Pokemon) PokemonAttacks {
	return PokemonAttacks{
		Fast:    toAttackList(pokemon.FastAttacks),
		Special: toAttackList(pokemon.SpecialAttacks),
	}
}

func toAttackList(attacks []repositories.PokemonAttack) []AttackDetails {
	var attackDetails []AttackDetails

	for _, attack := range attacks {
		attackDetails = append(attackDetails, AttackDetails{
			Name:   attack.Name,
			Type:   attack.Type,
			Damage: attack.Damage,
		})
	}

	return attackDetails
}

func toStringTypes(types []repositories.PokemonType) []string {
	var result []string

	for _, aType := range types {
		result = append(result, aType.Name)
	}

	return result
}
