package repositories

type Pokemon struct {
	Id             string
	Name           string
	Classification string
	Types          []PokemonType
	Resistant      []PokemonAttack
	Weaknesses     []PokemonAttack
}

type PokemonType struct {
	Name string
}

type PokemonAttack struct {
	Name string
}
