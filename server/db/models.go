package db

import "gorm.io/gorm"

type Tabler interface {
	TableName() string
}

type PokemonEntity struct {
	gorm.Model
	Identifier            string `gorm:"uniqueIndex"`
	Name                  string `gorm:"uniqueIndex"`
	Classification        string
	Types                 []PokemonEntityType       `gorm:"many2many:types_to_pokemon_type;"`
	Resistant             []PokemonEntityAttackType `gorm:"many2many:attacks_to_pokemon_resistant;"`
	Weaknesses            []PokemonEntityAttackType `gorm:"many2many:attacks_to_pokemon_weaknesses;"`
	MaxWeight             string
	MinWeight             string
	MaxHeight             string
	MinHeight             string
	FleeRate              float64
	EvolutionRequirements PokemonEntityEvolutionRequirements `gorm:"foreignKey:PokemonId"`
	Evolutions            []*PokemonEntity                   `gorm:"many2many:pokemon_evolutions"`
	PreviousEvolutions    []*PokemonEntity                   `gorm:"many2many:pokemon_previous_evolutions"`
	MaxCP                 int
	MaxHP                 int
	FastAttacks           []PokemonEntityAttack `gorm:"foreignKey:PokemonId"`
	SpecialAttacks        []PokemonEntityAttack `gorm:"foreignKey:PokemonId"`
}

func (PokemonEntity) TableName() string {
	return "pokemons"
}

type PokemonEntityType struct {
	gorm.Model
	Id   int
	Name string
}

func (PokemonEntityType) TableName() string {
	return "pokemon_types"
}

type PokemonEntityAttackType struct {
	gorm.Model
	Id   int
	Name string
}

func (PokemonEntityAttackType) TableName() string {
	return "attack_types"
}

type PokemonEntityEvolutionRequirements struct {
	gorm.Model
	Id        int
	Name      string
	Amount    int
	PokemonId int
}

func (PokemonEntityEvolutionRequirements) TableName() string {
	return "evolution_requirements"
}

type PokemonEntityAttack struct {
	gorm.Model
	Id        int
	Name      string
	TypeId    int
	Type      PokemonEntityAttackType `gorm:"foreignKey:TypeId"`
	Damage    int
	PokemonId int
}

func (PokemonEntityAttack) TableName() string {
	return "pokemon_attacks"
}
