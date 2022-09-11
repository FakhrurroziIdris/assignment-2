package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func RouteOrders(route *gin.Engine) {
	order := route.Group("/orders")
	order.GET("", controllers.GetOrders)
	order.POST("", controllers.CreateOrders)
	order.DELETE("/:orderID", controllers.DeleteOrders)
	order.PUT("/:orderID", controllers.UpdateOrders)
}
