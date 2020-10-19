package repositories

type Pokemon struct {
	Id             string
	Name           string
	Classification string
	Types          []PokemonType
}

type PokemonType struct {
	Name string
}
