package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/db"
	"gorm.io/gorm"
)

func Pokemons(dbConnection *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		pokemons := db.ListPokemons(dbConnection)
		c.JSON(200, pokemons)
	}
}
