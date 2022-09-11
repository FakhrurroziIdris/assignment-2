package routers

import (
	"assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func RouteCar(route *gin.Engine) {
	car := route.Group("/car")

	car.GET("", controllers.GetAll)
	car.GET("/:carID", controllers.GetCar)
	car.POST("", controllers.CreateCar)
	car.PUT("/:carID", controllers.UpdateCar)
	car.DELETE("/:carID", controllers.DeleteCar)
}
