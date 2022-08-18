package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zeralux/CryptoPayment/api/constant"
	"github.com/zeralux/CryptoPayment/api/controller/v1"
)

func GetApiRouter() *gin.Engine {
	// gin start
	router := gin.Default()

	// 對外API, GateWay不檢查token
	ver1Url := router.Group(constant.Ver1Url)
	{
		ver1Url.GET(constant.UsersUrl, v1.GetUsers)
		ver1Url.GET(constant.UsersFileUrl, v1.GetUsersFile)
	}

	// 對外API, GateWay檢查token
	privateUrl := router.Group(constant.PrivateUrl)
	{
		ver1Url := privateUrl.Group(constant.Ver1Url)
		{
			ver1Url.GET(constant.UserIdUrl, v1.GetUser)
		}
	}

	// 對內API, 內網即可調用
	internalUrl := router.Group(constant.InternalUrl)
	{
		ver1Url := internalUrl.Group(constant.Ver1Url)
		{
			ver1Url.POST(constant.UserUrl, v1.InsertUser)
			ver1Url.PATCH(constant.UserIdUrl, v1.UpdateUser)
			ver1Url.DELETE(constant.UserIdUrl, v1.DeleteUser)
		}
	}

	return router
}
