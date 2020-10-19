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

func toPokemonTypes(pokemonDefinitionTypes []string) []PokemonEntityType {
	var pokemonTypes []PokemonEntityType

	for _, pokemonType := range pokemonDefinitionTypes {
		pokemonTypes = append(pokemonTypes, PokemonEntityType{Name: pokemonType})
	}

	return pokemonTypes
}

func toEvolutions(pokemonDefinitionEvolutions []PokemonDefinitionEvolution) []Evolution {
	var evolutions []Evolution

	for _, pokemonDefinitionEvolution := range pokemonDefinitionEvolutions {
		evolutions = append(evolutions, Evolution{
			Id:   pokemonDefinitionEvolution.Id,
			Name: pokemonDefinitionEvolution.Name,
		})
	}

	return evolutions
}

func toPokemonAttacks(pokemonDefintion PokemonDefinition) PokemonAttacks {
	return PokemonAttacks{
		Fast:    toAttacks(pokemonDefintion.Attacks.Fast),
		Special: toAttacks(pokemonDefintion.Attacks.Special),
	}
}

func toAttacks(attackDefinitions []AttackDefinition) []Attack {
	var attacks []Attack

	for _, attackDefinition := range attackDefinitions {
		attacks = append(attacks, Attack{
			Name:   attackDefinition.Name,
			Type:   attackDefinition.Type,
			Damage: attackDefinition.Damage,
		})
	}

	return attacks
}
