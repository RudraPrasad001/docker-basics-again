package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{"Message": "Getting All users"})
}
func GetUser(c *gin.Context) {
	json := c.Query("username")
	if len(json) == 0 {
		message := "User does not exist"
		c.JSON(200, gin.H{"Message": message})
	} else {
		message := "Getting data for user " + json
		c.JSON(200, gin.H{"Message": message})
	}
}
