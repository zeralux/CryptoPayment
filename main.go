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

var servicePort = flag.Int("service.port", 8080, "http service port")
var configLocation = flag.String("config.location", "application-local.yml", "config file location")
var AppConfig *ApplicationYml

func main() {
	flag.Parse()
	// 讀取yml
	loadAppConfig()

	router := api.GetApiRouter()
	// swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//err := router.Run("0.0.0.0:" + strconv.Itoa(*servicePort))
	err := router.Run(":8080")
	if err != nil {
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
