package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/config"
	"github.com/ydarias/pokemon-go/server/controllers"
	"github.com/ydarias/pokemon-go/server/db"
	"github.com/ydarias/pokemon-go/server/repositories"
)

func setupRouter() *gin.Engine {
	dbConnection := db.Connect()

	populator := db.Populator{DbConnection: dbConnection}
	populator.Populate()

	pokemonRepository := repositories.PokemonDbRepository{DbConnection: dbConnection}
	pokemonTypesRepository := repositories.PokemonTypesDbRepository{DbConnection: dbConnection}

	pokemonController := controllers.PokemonsController{PokemonRepository: pokemonRepository}
	pokemonTypeController := controllers.PokemonTypesController{PokemonTypesRepository: pokemonTypesRepository}

	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.GET("/pokemons", pokemonController.Get)
	router.GET("/pokemons/types", pokemonTypeController.Get)
	router.GET("/pokemons/id/:pokemonId", pokemonController.GetById)
	router.GET("/pokemons/name/:pokemonName", pokemonController.GetByName)

	return router
}

func main() {
	port := config.Port()
	router := setupRouter()
	router.Run(":" + port)
}
