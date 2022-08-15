package main

import (
	"flag"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeralux/CryptoPayment/api"
	"github.com/zeralux/CryptoPayment/docs"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var configLocation = flag.String("config.location", "application-local.yml", "config file location")
var AppConfig *ApplicationYml

func main() {
	flag.Parse()
	// 讀取yml
	readAppConfig()

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	router := api.GetRouter()
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run()
}

type ApplicationYml struct {
	//Wallet struct {
	//	Limit struct {
	//		Address map[string]int64 `yaml:",flow"`
	//	} `yaml:"balance-limit"`
	//}
}

func readAppConfig() {
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
