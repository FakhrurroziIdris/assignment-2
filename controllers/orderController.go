package controllers

import (
	"net/http"
	"strconv"

	"assignment-2/models"
	service "assignment-2/service"

	"github.com/gin-gonic/gin"
)

// @BasePath /orders

// Get orders godoc
// @Summary get all orders
// @Schemes model.Order
// @Description get all orders
// @Tags example
// @Produce json
// @Success 200 {array} models.Order "OK"
// @Failure 404 {string} string "http.StatusNotFound"
// @Router / [get]
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

// Create orders
// @Summary create an order
// @Schemes model.Order
// @Description create an order
// @Tags example
// @Produce json
// @Success 201 {object} models.Order "OK"
// @Failure 400 {string} string "json string {message:{}}"
// @Failure 500 {string} string "json string {message:{}}"
// @Router / [post]
// @Param object body models.Order false "Order models"
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

// Delete order by ID
// @Summary delete an order
// @Schemes model.Order
// @Description delete an order
// @Tags example
// @Produce json
// @Success 200 {object} models.Order "OK"
// @Failure 400 {object} string "json string {message:{}}"
// @Failure 500 {string} string "json string {message:{}}"
// @Router /:orderID [delete]
// @Param uint64 query uint64 false "orderID"
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
	ctx.JSON(http.StatusOK, gin.H{
		"data": deleted,
	})
}

// Update order by ID
// @Summary update an order
// @Schemes model.Order
// @Description update an order
// @Tags example
// @Produce json
// @Success 200 {object} models.Order "OK"
// @Failure 400 {object} string "json string {message:{}}"
// @Failure 500 {string} string "json string {message:{}}"
// @Router /:orderID [put]
// @Param uint64 query uint64 false "orderID"
// @Param object body models.Order false "Order models"
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
		"data": deleted,
	})
}
