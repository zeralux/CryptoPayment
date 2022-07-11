package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controllerV1 "github.com/zeralux/gin/controller/v1"
	docs "github.com/zeralux/gin/docs"
)

func main() {
	router := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 對外API, GateWay不檢查token
	commonApi := router.Group("/")
	{
		v1 := commonApi.Group("/v1")
		{
			user := v1.Group("/users")
			{
				user.GET("", controllerV1.GetUsers)
			}
		}
	}

	// 對外API, GateWay檢查token
	publicApi := router.Group("/private")
	{
		v1 := publicApi.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.GET("/:id", controllerV1.GetUser)
			}
		}
	}

	// 對內API, 內網即可調用
	internalApi := router.Group("/internal")
	{
		v1 := internalApi.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("", controllerV1.InsertUser)
				user.PATCH("/:id", controllerV1.UpdateUser)
				user.DELETE("/:id", controllerV1.DeleteUser)
			}
		}
	}

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router.Run()
}
