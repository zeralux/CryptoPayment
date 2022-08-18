package main

import (
	"flag"
	"fmt"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeralux/CryptoPayment/api"
	"github.com/zeralux/CryptoPayment/docs"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var servicePort = flag.Int("service.port", 8080, "http service port")
var configLocation = flag.String("config.location", "application-local.yml", "config file location")
var AppConfig *ApplicationYml

func main() {
	flag.Parse()
	// 讀取yml
	loadAppConfig()
	// api router
	router := api.GetRouter()
	// swagger router
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// run service
	if err := router.Run(fmt.Sprintf(":%d", *servicePort)); err != nil {
		panic(err)
	}
}

type ApplicationYml struct {
	//Wallet struct {
	//	Limit struct {
	//		Address map[string]int64 `yaml:",flow"`
	//	} `yaml:"balance-limit"`
	//}
}

func loadAppConfig() {
	var err error
	var bytes []byte
	if bytes, err = ioutil.ReadFile(*configLocation); err != nil {
		panic(err)
	}
	AppConfig = &ApplicationYml{}
	if err = yaml.Unmarshal(bytes, AppConfig); err != nil {
		panic(err)
	}
}
