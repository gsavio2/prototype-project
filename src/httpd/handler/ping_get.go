package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping1 test ###
func Ping1(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

// Ping2 test ###
func Ping2(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"message": "pinga",
	})
}

// Ping3 test ###
func Ping3(k, v string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			k: v,
		})
	}
}
