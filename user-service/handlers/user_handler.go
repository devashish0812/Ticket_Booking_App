package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", GetUsers)
		userGroup.GET("/:id", GetUserByID)
		userGroup.POST("/", CreateUser)
		userGroup.PUT("/:id", UpdateUser)
		userGroup.DELETE("/:id", DeleteUser)
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "List of users"})
}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get user by ID", "id": id})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User updated", "id": id})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "id": id})
}
