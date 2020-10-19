package controllers

type PokemonView struct {
	Id             string
	Name           string
	Classification string
	Types          []string
	Resistant      []string
	Weaknesses     []string
	Weight         RangeView
	Height         RangeView
	FleeRate       float64
	MaxCP          int
	MaxHP          int
}

type RangeView struct {
	Maximum string
	Minimum string
}
