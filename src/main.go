package main

import (
	"github.com/gin-gonic/contrib/cors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// rotas /api

	api := r.Group("/api")
	{
		// get ping
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
		api.GET("")
	}

	r.Use(cors.Default())
	r.Run()
}
