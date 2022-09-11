package controllers

import (
	"assignment-2/models"
	"fmt"
	"net/http"
	"strconv"

	brandService "assignment-2/service"

	"github.com/gin-gonic/gin"
)

func CreateBrand(ctx *gin.Context) {
	newBrand := models.Brand{}
	if err := ctx.ShouldBindJSON(&newBrand); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	created, err := brandService.CreateBrand(newBrand)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"car": created,
	})
}

func UpdateBrand(ctx *gin.Context) {
	brandID, err := strconv.ParseUint(ctx.Param("brandID"), 0, 0)
	var updatedBrand models.Brand

	if err := ctx.ShouldBindJSON(&updatedBrand); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBrand.ID = brandID
	updated, affected, err := brandService.UpdateBrand(updatedBrand)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  err,
			"error_message": err,
		})
		return
	} else if affected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Brand with id %v not found", brandID),
			"status":        http.StatusNotModified,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Brand successfully updated to %v", updated),
		"data":    updated,
	})
}

func GetBrand(ctx *gin.Context) {
	brandID, err := strconv.ParseUint(ctx.Param("brandID"), 0, 0)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	brand, err := brandService.GetBrand(brandID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Brand with id %v not found", brandID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"brand": brand,
	})
}

// func DeleteCar(ctx *gin.Context) {
// 	carID := ctx.Param("carID")
// 	condition := false
// 	var car models.Car
// 	var carIndex int

// 	if err := ctx.ShouldBindJSON(&car); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	for i, car := range models.Cars {
// 		if carID == car.CarID {
// 			condition = true
// 			carIndex = i
// 			break
// 		}
// 	}

// 	if !condition {
// 		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
// 			"error_status":  "Data not found",
// 			"error_message": fmt.Sprintf("car with id %v not found", carID),
// 		})
// 		return
// 	}

// 	copy(models.Cars[carIndex:], models.Cars[carIndex+1:])
// 	models.Cars[len(models.Cars)-1] = models.Car{}
// 	models.Cars = models.Cars[:len(models.Cars)-1]

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message": fmt.Sprintf("car with id %v has been successfully deleted", carID),
// 	})
// }

// func GetAll(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"All cars": models.Cars,
// 	})
// }
