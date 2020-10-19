package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/config"
	"github.com/ydarias/pokemon-go/server/controllers"
	"github.com/ydarias/pokemon-go/server/db"
)

func main() {
	dbConnection := db.Connect()

	db.Populate(dbConnection)

	port := config.Port()

	router := gin.Default()

	router.GET("/ping", controllers.Ping)
	router.GET("/pokemons", controllers.Pokemons(dbConnection))

	router.Run(":" + port)
}
