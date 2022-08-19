package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zeralux/CryptoPayment/api/model"
	"io"
	"net/http"
	"os"
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
// @param level query string false "職級" Enums(A,B,C)
// @param method query []string false "支付方式" Enums(CreditCard,LinePay,ApplePay)
// @Produce json
// @Router /v1/users [get]
func GetUsers(c *gin.Context) {
	name := c.Request.URL.Query().Get("name")
	level := c.Request.URL.Query().Get("level")
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"level": level,
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

// GetUsers
// @Tags 用戶
// @Summary 下載
// @param file path string true "檔案類型" default(csv)
// @Produce application/octet-stream
// @Router /v1/users/{file} [get]
func GetUsersFile(c *gin.Context) {
	downloadName := "test.csv"
	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + downloadName}
	file, err := os.Open("test.csv")
	if err != nil {
		c.String(http.StatusOK, "%v", err)
		return
	}
	defer file.Close()
	io.Copy(c.Writer, file)
}
