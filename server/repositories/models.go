package repositories

type Pokemon struct {
	Id             string
	Name           string
	Classification string
	Types          []PokemonType
	Resistant      []PokemonAttack
	Weaknesses     []PokemonAttack
	MaxWeight      string
	MinWeight      string
	MaxHeight      string
	MinHeight      string
	FleeRate       float64
	MaxCP          int
	MaxHP          int
}

type PokemonType struct {
	Name string
}

type PokemonAttack struct {
	Name string
}
