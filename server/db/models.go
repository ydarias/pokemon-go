package db

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

type PokemonEntity struct {
	gorm.Model
	Identifier     string
	Name           string
	Classification string
	Types          []PokemonEntityType   `gorm:"many2many:types_to_pokemon_type;"`
	Resistant      []PokemonEntityAttack `gorm:"many2many:attacks_to_pokemon_resistant;"`
	Weaknesses     []PokemonEntityAttack `gorm:"many2many:attacks_to_pokemon_weaknesses;"`
	MaxWeight      string
	MinWeight      string
	MaxHeight      string
	MinHeight      string
	FleeRate       float64
	//		EvolutionRequirements struct {
	//			Amount int
	//			Name   string
	//		}
	//		Evolutions []Evolution
	MaxCP int
	MaxHP int
	// Attacks PokemonAttacks
}

func (PokemonEntity) TableName() string {
	return "Pokemons"
}

type PokemonEntityType struct {
	gorm.Model
	Name string
}

type PokemonEntityAttack struct {
	gorm.Model
	Name string
}

func (PokemonEntityType) TableName() string {
	return "PokemonTypes"
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
