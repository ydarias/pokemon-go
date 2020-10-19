package controllers

type PokemonView struct {
	Id             string
	Name           string
	Classification string
	Types          []string
	Resistant      []string
	Weaknesses     []string
}
