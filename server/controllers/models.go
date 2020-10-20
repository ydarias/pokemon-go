package controllers

type PokemonView struct {
	Id                    string
	Name                  string
	Classification        string
	Types                 []string
	Resistant             []string
	Weaknesses            []string
	Weight                RangeView
	Height                RangeView
	FleeRate              float64
	EvolutionRequirements EvolutionRequirements
	MaxCP                 int
	MaxHP                 int
	Evolutions            []EvolutionReference
	PreviousEvolutions    []EvolutionReference
	Attacks               PokemonAttacks
}

type RangeView struct {
	Maximum string
	Minimum string
}

type EvolutionRequirements struct {
	Amount int
	Name   string
}

type EvolutionReference struct {
	Id   int
	Name string
}

type PokemonAttacks struct {
	Fast    []AttackDetails
	Special []AttackDetails
}

type AttackDetails struct {
	Name   string
	Type   string
	Damage int
}
