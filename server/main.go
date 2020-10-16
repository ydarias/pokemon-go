package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydarias/pokemon-go/server/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/ping", controllers.Ping)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
