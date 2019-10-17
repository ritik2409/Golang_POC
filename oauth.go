package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ginoauth2 "github.com/zalando/gin-oauth2"
	"github.com/zalando/gin-oauth2/zalando"
)

func main() {
	var USERS []zalando.AccessTuple = []zalando.AccessTuple{
		{"/employees", "sszuecs", "Sandor Szücs"},
		{"/employees", "njuettner", "Nick Jüttner"},
	}

	router := gin.Default()
	// router.Use(ginglog.Logger(3 * time.Second))
	// router.Use(ginoauth2.RequestLogger([]string{"uid"}, "data"))
	// router.Use(gin.Recovery())

	privateUser := router.Group("/api/privateUser")
	privateUser.Use(ginoauth2.Auth(zalando.UidCheck(USERS), zalando.OAuth2Endpoint))
	privateUser.GET("/", func(c *gin.Context) {
		if v, ok := c.Get("cn"); ok {
			c.JSON(200, gin.H{"message": fmt.Sprintf("Hello from private for users to %s", v)})
		} else {
			c.JSON(200, gin.H{"message": "Hello from private for users without cn"})
		}
	})

	router.Run(":8080")
}
