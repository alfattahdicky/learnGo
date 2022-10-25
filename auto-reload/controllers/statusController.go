package controllers

import (
	"auto-reload/models"
	"fmt"
	"math/rand"
	"net/http"
	_ "time"

	"github.com/gin-gonic/gin"
)

func GetAllStatus(ctx *gin.Context)  {
	newStatus := models.Status{
		Water: rand.Intn(100),
		Wind: rand.Intn(100),
	}

	fmt.Println(newStatus)

	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.JSON(http.StatusOK, gin.H{
		"status": newStatus,
	})
}