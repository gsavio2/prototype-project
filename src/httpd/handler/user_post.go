package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrohmachado/prototype-project/src/platform/user"
)

//UserPostRequest struct
type UserPostRequest struct {
	//attr type tag
	Name  string `json: "name"`
	Email string `json: "email"`
}

//UserPost post user
func UserPost(users user.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UserPostRequest{}
		c.Bind(&requestBody)

		user := user.User{
			Name:  requestBody.Name,
			Email: requestBody.Email,
		}

		users.Add(user)

		c.Status(http.StatusNoContent)
	}
}
