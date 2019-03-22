package main

import (
	"github.com/gin-contrib/cors"
	"github.com/pedrohmachado/prototype-project/src/httpd/handler"
	"github.com/pedrohmachado/prototype-project/src/platform/user"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowAllOrigins permite requisição de qualquer origem (depois configurar corretamente)
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           86400,
	}))

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
		api.PUT("/user/{name}", handler.UserPut(users))
	}

	r.Run(":8081")
}
