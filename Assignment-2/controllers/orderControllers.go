package controllers

import (
	"assignment-two/database"
	"assignment-two/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context)  {
	db := database.GetDb()

	var Order models.Orders

	ctx.ShouldBindJSON(&Order)

	err := db.Create(&Order).Error

	if err != nil {
		fmt.Println("Error creating user data", err)

		return
	}

	ctx.JSON(http.StatusOK, Order)

	fmt.Println("New User Data", Order)
}

func GetAllOrdersWithItems(ctx *gin.Context)  {
	db := database.GetDb()
	var orders []models.Orders

	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Cannot get all orders", err)
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	fmt.Println("Success Get All orders", orders)

	ctx.JSON(http.StatusOK, orders)
}

func DeleteOrderById(ctx *gin.Context)  {
	db := database.GetDb()

	Id,e := strconv.Atoi(ctx.Param("orderId")) 

	if e != nil {
		fmt.Println("Must be integer", e)

		return
	}

	Order := models.Orders{}

	err := db.Where("id = ?", Id).Delete(&Order).Error

	if err != nil {
		fmt.Println("Error Deleting Order: ", err.Error())

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not found",
			"error_message": fmt.Sprintf("order with %d not found", Id),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"message": fmt.Sprintf("order with id %d has been successfully deleted", Id), 
	})
}

func UpdateOrderById(ctx *gin.Context)  {
	db := database.GetDb()
	Id,e := strconv.Atoi(ctx.Param("orderId"))

	if e != nil {
		fmt.Println("Id must be integer",e)
		return
	}

	OrderUpdate := models.Orders{}
	
	ctx.ShouldBindJSON(&OrderUpdate)

	err := db.Model(&OrderUpdate).Where("id = ?", Id).Updates(&OrderUpdate)

	// err := db.Exec("")

	if err != nil {
		fmt.Println("Error Updating Order By Id:", err.Error)

		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "Data Not found",
			"error_message": fmt.Sprintf("Order with id %d not found", Id),
		})

		return
	}

	fmt.Println("Success Updating Order By Id", OrderUpdate)

	ctx.JSON(http.StatusOK, OrderUpdate)
}