package main

import (
	"assignment-2/configs"
	pgsql "assignment-2/helpers/database"

	routerConfig "assignment-2/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config := configs.InitConfig()
	pgsql.StartDB(config.Database)

	// brand.Name = "Updated Toyota"
	// brandRepo.UpdateBrand(brand)

	// brandRepo.DeleteBrand(brand)

	var PORT = config.WebServer.Port

	router := gin.Default()
	routerConfig.RouteOrders(router)
	// routerConfig.RouteCar(router)
	// routerConfig.RouteBrand(router)
	router.Run(PORT)
}
