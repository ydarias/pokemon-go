package repositories

import (
	"github.com/go-playground/assert/v2"
	"github.com/ydarias/pokemon-go/server/db"
	"gorm.io/driver/sqlite"
	"testing"
)

func TestPokemonRepository(t *testing.T) {
	connection := db.Connection{
		Dialector: sqlite.Open("file::memory:?cache=shared"),
	}
	dbConnection := connection.Connect()

	populator := db.Populator{DbConnection: dbConnection}
	populator.Populate()

	pokemonRepository := PokemonDbRepository{DbConnection: dbConnection}

	t.Run("puturru", func(t *testing.T) {
		pokemons := pokemonRepository.FindPokemons()

		assert.Equal(t, 3, len(pokemons))
	})

	t.Run("should recover a pokemon by id", func(t *testing.T) {
		pokemon := pokemonRepository.FindPokemonById("001")

		assert.Equal(t, "Bulbasaur", pokemon.Name)
	})

	t.Run("should recover a pokemon by id", func(t *testing.T) {
		pokemon := pokemonRepository.FindPokemonByName("Ivysaur")

		assert.Equal(t, "002", pokemon.Id)
	})

	underliyingDb, _ := dbConnection.DB()
	underliyingDb.Close()
}
