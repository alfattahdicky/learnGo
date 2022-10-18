package controllers

import (
	"auto-reload/models"
	"net/http"
	"time"
	"math/rand"
	"github.com/gin-gonic/gin"
)


func GetAllStatus(ctx *gin.Context)  {
	var waterRand int
	var windRand int
	for{
		time.Sleep(15* time.Second)
		waterRand = rand.Intn(100)
		windRand = rand.Intn(100)
	}
	var newStatus models.Status = models.Status{
		Water: waterRand,
		Wind: windRand,
	}
	err := ctx.ShouldBindJSON(&newStatus)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": newStatus,
	})
}