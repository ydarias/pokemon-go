package repositories

import (
	"fmt"
	"github.com/ydarias/pokemon-go/server/db"
	"gorm.io/gorm"
)

type PokemonRepository interface {
	FindPokemons() []Pokemon
}

type PokemonDbRepository struct {
	DbConnection *gorm.DB
}

func (p PokemonDbRepository) FindPokemons() []Pokemon {
	var pokemonEntities []db.PokemonEntity
	result := p.DbConnection.Preload("Types").Preload("Resistant").Preload("Weaknesses").Find(&pokemonEntities)
	if result.Error != nil {
		fmt.Print(result.Error)
	}

	return toPokemons(pokemonEntities)
}

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
