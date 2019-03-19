package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}

func main() {

	r := gin.Default()

	r.GET("/", helloWorld)

	r.Run(":8000")
}
