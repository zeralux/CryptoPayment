package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser " + id,
	})
}

func GetUsers(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser " + id,
	})
}

func InsertUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "InsertUser",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser",
	})
}
