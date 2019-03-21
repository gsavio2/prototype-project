package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrohmachado/prototype-project/src/platform/user"
)

//UserGet get all users
func UserGet(users user.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := users.GetAll()
		c.JSON(http.StatusOK, results)
	}
}
