package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type PokemonDefinition struct {
	Id             string
	Name           string
	Classification string
	Types          []string
	Resistant      []string
	Weaknesses     []string
	Weight         struct {
		Minimum string
		Maximum string
	}
	Height struct {
		Minimum string
		Maximum string
	}
	FleeRate              float64
	EvolutionRequirements struct {
		Amount int
		Name   string
	}
	Evolutions         []PokemonDefinitionEvolution
	PreviousEvolutions []PokemonDefinitionEvolution
	MaxCP              int
	MaxHP              int
	Attacks            struct {
		Fast    []AttackDefinition
		Special []AttackDefinition
	}
}

type PokemonDefinitionEvolution struct {
	Id   int
	Name string
}

type AttackDefinition struct {
	Name   string
	Type   string
	Damage int
}

func ParsePokemons() []PokemonDefinition {
	data, err := ioutil.ReadFile("./pokemons.json")
	if err != nil {
		fmt.Print(err)
	}

	var pokemons []PokemonDefinition

	err = json.Unmarshal(data, &pokemons)
	if err != nil {
		fmt.Println("error:", err)
	}

	return pokemons
}
