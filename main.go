package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeralux/gin/constant"
	controllerV1 "github.com/zeralux/gin/controller/v1"
	docs "github.com/zeralux/gin/docs"
)

func main() {
	router := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 對外API, GateWay不檢查token
	ver1Url := router.Group(constant.Ver1Url)
	{
		ver1Url.GET(constant.UsersUrl, controllerV1.GetUsers)
		ver1Url.GET(constant.UsersFileUrl, controllerV1.GetUsersFile)
	}

	// 對外API, GateWay檢查token
	privateUrl := router.Group(constant.PrivateUrl)
	{
		ver1Url := privateUrl.Group(constant.Ver1Url)
		{
			ver1Url.GET(constant.UserIdUrl, controllerV1.GetUser)
		}
	}

	// 對內API, 內網即可調用
	internalUrl := router.Group(constant.InternalUrl)
	{
		ver1Url := internalUrl.Group(constant.Ver1Url)
		{
			ver1Url.POST(constant.UserUrl, controllerV1.InsertUser)
			ver1Url.PATCH(constant.UserIdUrl, controllerV1.UpdateUser)
			ver1Url.DELETE(constant.UserIdUrl, controllerV1.DeleteUser)
		}
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
