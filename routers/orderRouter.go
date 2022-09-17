package routers

import (
	"assignment-2/controllers"
	"fmt"

	"github.com/gin-gonic/gin"

	docs "assignment-2/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Orders API
// @version 1.0
// @description This is a sample service for managing orders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fakhrurrozi.idris@gmail.com
// @license no
// @host localhost:8080
// @BasePath /orders
func RouteOrders(route *gin.Engine) {
	order := route.Group("/orders")
	order.GET("", controllers.GetOrders)
	order.POST("", controllers.CreateOrders)
	order.DELETE("/:orderID", controllers.DeleteOrders)
	order.PUT("/:orderID", controllers.UpdateOrders)

	docs.SwaggerInfo.BasePath = "/orders"
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/orders/doc.json", ":8080")) // The url pointing to API definition
	route.GET("/swagger/orders/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
