package main

import (
	"assignment-2/configs"
	routerConfig "assignment-2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.InitConfig()
	var PORT = config.WebServer.Port

	router := gin.Default()
	routerConfig.RouteOrders(router)
	// routerConfig.RouteCar(router)
	// routerConfig.RouteBrand(router)

	// docs.SwaggerInfo.BasePath = "/orders"
	// url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", PORT)) // The url pointing to API definition
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run(PORT)
}
