package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeralux/gin/model"
	"net/http"
)

// GetUser
// @Tags 用戶
// @Summary 查單筆
// @param id path string true "會員id"
// @Produce json
// @Router /private/v1/user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUser id=" + id,
	})
}

// GetUsers
// @Tags 用戶
// @Summary 查多筆
// @param name query string true "會員姓名"
// @Produce json
// @Router /v1/users [get]
func GetUsers(c *gin.Context) {
	name := c.Request.URL.Query().Get("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUsers name=" + name,
	})
}

// InsertUser
// @Tags 用戶
// @Summary 新增
// @param user body model.User true "會員資料"
// @Produce json
// @Router /internal/v1/user [post]
func InsertUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	body, _ := json.Marshal(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "InsertUser " + string(body),
	})
}

// UpdateUser
// @Summary 部份更新
// @Schemes
// @Description 說明
// @Tags 用戶
// @Accept json
// @Produce json
// @Router /internal/v1/user/{id} [patch]
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
// @Router /internal/v1/user/{id} [delete]
func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UpdateUser",
	})
}
