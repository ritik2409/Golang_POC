package main

import (
	"project/db"
	"project/migrations"
	"project/server"
	//         // "github.com/gin-gonic/gin"
	// 		// "strconv"
	//         // "fmt"
)

func main() {
	// // r := versioning(gin.Default())
	// router := gin.Default()
	// v1 := router.Group("api/v1")
	// v2 :=  router.Group("api/v2")
	// // v1 := router.Group(getversionname())
	// {
	//     v.POST("/users", PostUser)
	//     v.GET("/users", GetUsers)
	//     v.GET("/users/:id", GetUser)
	//     v.PUT("/users/:id", UpdateUser)
	//     v.DELETE("/users/:id", DeleteUser)
	// }
	// router.Run(getportname()) // listen and serve on 0.0.0.0:8080
	migrations.Gorm_migrations()
	db.Init()
	server.Init()

}
