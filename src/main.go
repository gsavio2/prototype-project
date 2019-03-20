package main

import (
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
}

func carregaUsuarios() {

}
