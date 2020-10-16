package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/config"
	"github.com/ydarias/pokemon-go/server/controllers"
)

func main() {
	port := config.Port()

	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	router.Run(":" + port)
}
