package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pedrohmachado/prototype-project/src/platform/user"
)

//UserPutRequest struct
type UserPutRequest struct {
	//attr type tag
	Name  string `json: "name"`
	Email string `json: "email"`
}

func UserPut(users user.Updater) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UserPutRequest{}
		c.Bind(&requestBody)

		user := user.User{
			Name:  requestBody.Name,
			Email: requestBody.Email,
		}

		name := c.Param("name")
		email := c.Param("email")

		users.Update(user, name)
	}
}
