package main

import (
	"github.com/pedrohmachado/prototype-project/src/httpd/handler"
	"github.com/pedrohmachado/prototype-project/src/platform/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	users := user.New()

	api := r.Group("/api")
	{
		ping := api.Group("/ping")
		{
			ping.GET("/one", handler.Ping1)
			ping.GET("/two", handler.Ping2)
			ping.GET("/three", handler.Ping3("message", "ping"))
		}

		api.GET("/user", handler.UserGet(users))
		api.POST("/user", handler.UserPost(users))
	}

	r.Run(":8081")
}
