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
	if !populator.IsPopulated() {
		populator.Populate()
	}

	pokemonRepository := repositories.PokemonDbRepository{DbConnection: dbConnection}

	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.GET("/pokemons", controllers.Pokemons(pokemonRepository))

	return router
}

func main() {
	port := config.Port()
	router := setupRouter()
	router.Run(":" + port)
}
