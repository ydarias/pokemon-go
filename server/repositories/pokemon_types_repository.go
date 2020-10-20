package repositories

import "gorm.io/gorm"

type PokemonTypesRepository interface {
	FindTypes() []PokemonType
}

type PokemonTypesDbRepository struct {
	DbConnection *gorm.DB
}

func (p PokemonTypesDbRepository) FindTypes() []PokemonType {
	var pokemonTypes []PokemonType
	p.DbConnection.Find(&pokemonTypes)

	return pokemonTypes
}
