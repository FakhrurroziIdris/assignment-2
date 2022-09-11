package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCar(ctx *gin.Context) {
	var newCar models.Car

	if err := ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	newCar.CarID = fmt.Sprintf("c%d", len(models.Cars)+1)
	models.Cars = append(models.Cars, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var updatedCar models.Car

	if err := ctx.ShouldBindJSON(&updatedCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range models.Cars {
		if carID == car.CarID {
			condition = true
			models.Cars[i] = updatedCar
			models.Cars[i].CarID = carID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carID),
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var car models.Car

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for k, v := range models.Cars {
		fmt.Println("CarID:", car.CarID)
		if carID == v.CarID {
			condition = true
			car = models.Cars[k]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": car,
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("carID")
	condition := false
	var car models.Car
	var carIndex int

	if err := ctx.ShouldBindJSON(&car); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, car := range models.Cars {
		if carID == car.CarID {
			condition = true
			carIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("car with id %v not found", carID),
		})
		return
	}

	copy(models.Cars[carIndex:], models.Cars[carIndex+1:])
	models.Cars[len(models.Cars)-1] = models.Car{}
	models.Cars = models.Cars[:len(models.Cars)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
	})
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"All cars": models.Cars,
	})
}
