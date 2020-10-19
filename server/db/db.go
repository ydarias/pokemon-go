package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&PokemonEntity{}, &PokemonEntityType{})
	if err != nil {
		panic("failed to apply schema changes")
	}

	return db
}

type Populator struct {
	DbConnection *gorm.DB
}

func (p Populator) IsPopulated() bool {
	var count int64
	p.DbConnection.Model(&PokemonEntity{}).Count(&count)

	return count > 0
}

func (p Populator) Populate() {
	pokemons := ParsePokemons()
	for _, pokemon := range pokemons {
		pokemonInstance := ToPokemonInstance(pokemon)
		pokemonInstance.Types = p.populateTypes(pokemon.Types)
		p.DbConnection.Where(PokemonEntity{Identifier: pokemon.Id}).FirstOrCreate(&pokemonInstance)
	}
}

func (p Populator) populateTypes(pokemonTypes []string) []PokemonEntityType {
	var pokemonEntityTypes []PokemonEntityType
	for _, pokemonType := range pokemonTypes {
		pokemonEntityTypes = append(pokemonEntityTypes, p.findOrCreateType(pokemonType))
	}

	return pokemonEntityTypes
}

func (p Populator) findOrCreateType(pokemonType string) PokemonEntityType {
	var pokemonEntityType PokemonEntityType
	p.DbConnection.Where(PokemonEntityType{Name: pokemonType}).FirstOrCreate(&pokemonEntityType)

	return pokemonEntityType
}
