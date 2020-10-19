package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Pokemon struct {
	gorm.Model
	Identifier     string
	Name           string
	Classification string
	Types          []PokemonType `gorm:"many2many:types_to_pokemon_type;"`
	/*

			Resistant []string
		Weaknesses []string
		Weight struct {
			Minimum string
			Maximum string
		}
		Height struct {
			Minimum string
			Maximum string
		}
		FleeRate float64
		EvolutionRequirements struct {
			Amount int
			Name string
		}
		Evolutions []Evolution
		MaxCP int
		MaxHP int
		Attacks PokemonAttacks

	*/
}

type PokemonType struct {
	gorm.Model
	Name string
}

type Evolution struct {
	gorm.Model
	Id   int
	Name string
}

type Attack struct {
	gorm.Model
	Name   string
	Type   string
	Damage int
}

type PokemonAttacks struct {
	gorm.Model
	Fast    []Attack
	Special []Attack
}

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Populate(dbConnection *gorm.DB) {
	err := dbConnection.AutoMigrate(&Pokemon{})
	if err != nil {
		panic("failed to apply schema changes")
	}

	pokemons := ParsePokemons()
	for _, pokemon := range pokemons {
		pokemonInstance := ToPokemonInstance(pokemon)
		dbConnection.Create(&pokemonInstance)
	}
}

func ListPokemons(dbConnection *gorm.DB) []Pokemon {
	var pokemons []Pokemon
	result := dbConnection.Find(&pokemons)
	if result.Error != nil {
		fmt.Print(result.Error)
	}

	return pokemons
}
