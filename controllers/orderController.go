package controllers

import (
	"net/http"
	"strconv"

	"assignment-2/models"
	service "assignment-2/service"

	"github.com/gin-gonic/gin"
)

func GetOrders(ctx *gin.Context) {
	orders, err := service.GetOrders()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": orders,
	})
}

func CreateOrders(ctx *gin.Context) {
	payload := models.Order{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	created, err := service.CreateOrders(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": created,
	})
}

func DeleteOrders(ctx *gin.Context) {
	var orderID uint64
	var err error

	if orderID, err = strconv.ParseUint(ctx.Param("orderID"), 0, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	deleted, err := service.DeleteOrders(orderID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"data": deleted,
	})
}

func UpdateOrders(ctx *gin.Context) {
	var orderID uint64
	var err error

	if orderID, err = strconv.ParseUint(ctx.Param("orderID"), 0, 0); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	payload := models.Order{}
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	payload.Order_ID = orderID
	deleted, err := service.UpdateOrders(payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"orders": deleted,
	})
}
