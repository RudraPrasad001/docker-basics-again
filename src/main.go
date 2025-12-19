package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudraprasad001/dockercompose/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterUserRoutes(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"Message": "Welcome to gin"})
	})

	r.Run(":8081")
}
