package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rudraprasad001/dockercompose/controllers"
)

func RegisterUserRoutes(r *gin.Engine) {
	group := r.Group("/users")
	group.GET("/", controllers.GetUsers)
	group.GET("/get", controllers.GetUser)

}
