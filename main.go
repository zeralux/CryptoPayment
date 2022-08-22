package main

import (
	"flag"
	"fmt"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeralux/CryptoPayment/api"
	"github.com/zeralux/CryptoPayment/config"
	"github.com/zeralux/CryptoPayment/docs"
)

var configLocation = flag.String("config.location", "application-local.yml", "config file location")

func main() {
	// get setting
	flag.Parse()
	config.LoadAppConfig(*configLocation)
	appConfig := config.GetAppConfig()

	// setting db

	// setting http
	router := api.GetRouter()
	// swagger router
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run service
	if err := router.Run(fmt.Sprintf(":%d", appConfig.Service().Port())); err != nil {
		panic(err)
	}
}
