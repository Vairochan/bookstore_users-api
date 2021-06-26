package app

import (
	"main/controllers/ping"
	"main/controllers/user"
)

func mapUrls(){
	router.GET("/ping", ping.Ping)

	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.GET("/users/search", user.Search)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)
	router.GET("/internal/users/search", user.Search)
}