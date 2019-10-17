package server

import (
	"project/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("api/v1")
	v2 := router.Group("api/v2")

	v1Controller := new(controller.Version1Controller)
	// v1Controller.Init()
	v2Controller := new(controller.Version2Controller)
	// v2Controller.Init()

	{
		v1.POST("/users", v1Controller.PostUser)
		v1.GET("/users", v1Controller.GetUsers)
		v1.GET("/users/:id", v1Controller.GetUser)
		v1.PUT("/users/:id", v1Controller.UpdateUser)
		v1.DELETE("/users/:id", v1Controller.DeleteUser)
	}
	{
		v2.POST("/users", v2Controller.PostUser)
		v2.GET("/users", v2Controller.GetUsers)
		v2.GET("/users/:id", v2Controller.GetUser)
		v2.PUT("/users/:id", v2Controller.UpdateUser)
		v2.DELETE("/users/:id", v2Controller.DeleteUser)
	}

	return router
}
