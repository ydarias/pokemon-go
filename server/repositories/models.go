package repositories

type Pokemon struct {
	Id                    string
	Name                  string
	Classification        string
	Types                 []PokemonType
	Resistant             []PokemonAttackType
	Weaknesses            []PokemonAttackType
	MaxWeight             string
	MinWeight             string
	MaxHeight             string
	MinHeight             string
	FleeRate              float64
	EvolutionRequirements EvolutionRequirements
	Evolutions            []EvolutionReference
	PreviousEvolutions    []EvolutionReference
	MaxCP                 int
	MaxHP                 int
	FastAttacks           []PokemonAttack
	SpecialAttacks        []PokemonAttack
}

type PokemonType struct {
	Name string
}

type PokemonAttackType struct {
	Name string
}

type EvolutionRequirements struct {
	Amount int
	Name   string
}

type EvolutionReference struct {
	Id   string
	Name string
}

type PokemonAttack struct {
	Name   string
	Type   string
	Damage int
}
