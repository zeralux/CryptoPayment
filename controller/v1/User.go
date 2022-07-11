package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUser
// @Summary 查單筆
// @Schemes
// @Description 說明
// @Tags 用戶
// @Produce json
// @Router /v1/user/:id [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser " + id,
	})
}

// GetUsers
// @Summary 查多筆
// @Schemes
// @Description 說明
// @Tags 用戶
// @Produce json
// @Router /v1/users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUsers",
	})
}

// InsertUser
// @Summary 新增
// @Schemes
// @Description 說明
// @Tags 用戶
// @Accept json
// @Produce json
// @Router /v1/user [post]
func InsertUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "InsertUser",
	})
}

// UpdateUser
// @Summary 部份更新
// @Schemes
// @Description 說明
// @Tags 用戶
// @Accept json
// @Produce json
// @Router /v1/user/:id [patch]
func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser",
	})
}

// DeleteUser
// @Summary 刪除
// @Schemes
// @Description 說明
// @Tags 用戶
// @Produce json
// @Router /v1/user/:id [delete]
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser",
	})
}
