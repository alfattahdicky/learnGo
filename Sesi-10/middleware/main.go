package main

import (
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func middleware(ctx *gin.Context) {
	auth := ctx.Query("auth")

	if auth != "12345" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H {
			"message": "unauthorized",
		})

		return
	}

	ctx.Next()
}

func middleware2(ctx *gin.Context)  {
	ctx.Next()
}

func main() {
	r := gin.Default()

	// middleware /intercepted
	// r.Use(middleware, middleware2)
	r.Use(middleware)
	r.Use(middleware2)


	r.GET("/ping", func (c *gin.Context)  {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
	r.Run()
}