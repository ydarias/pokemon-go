package db

func ToPokemonEntity(pokemon PokemonDefinition) PokemonEntity {
	return PokemonEntity{
		Identifier:     pokemon.Id,
		Name:           pokemon.Name,
		Classification: pokemon.Classification,
		MaxWeight:      pokemon.Weight.Maximum,
		MinWeight:      pokemon.Weight.Minimum,
		MaxHeight:      pokemon.Height.Maximum,
		MinHeight:      pokemon.Height.Minimum,
		FleeRate:       pokemon.FleeRate,
		MaxCP:          pokemon.MaxCP,
		MaxHP:          pokemon.MaxHP,
	}
}
