package repositories

import (
	"fmt"
	"github.com/ydarias/pokemon-go/server/db"
	"gorm.io/gorm"
)

type PokemonRepository interface {
	FindPokemons() []Pokemon
	FindPokemonById(id string) Pokemon
	FindPokemonByName(name string) Pokemon
}

type PokemonDbRepository struct {
	DbConnection *gorm.DB
}

func (p PokemonDbRepository) FindPokemons() []Pokemon {
	var pokemonEntities []db.PokemonEntity
	result := p.DbConnection.
		Preload("Types").
		Preload("Resistant").
		Preload("Weaknesses").
		Preload("EvolutionRequirements").
		Preload("Evolutions").
		Preload("PreviousEvolutions").
		Preload("FastAttacks").
		Preload("FastAttacks.Type").
		Preload("SpecialAttacks").
		Preload("SpecialAttacks.Type").
		Find(&pokemonEntities)
	if result.Error != nil {
		fmt.Print(result.Error)
	}

	return toPokemons(pokemonEntities)
}

func (p PokemonDbRepository) FindPokemonById(id string) Pokemon {
	return Pokemon{}
}

func (p PokemonDbRepository) FindPokemonByName(name string) Pokemon {
	return Pokemon{}
}
