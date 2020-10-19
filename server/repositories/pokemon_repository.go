package repositories

type PokemonRepository interface {
	findPokemons() []Pokemon
}
