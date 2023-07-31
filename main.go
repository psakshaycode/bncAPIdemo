package main

import (
	"bncapidemo/handlers"
	"bncapidemo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/assets", func(ctx *gin.Context) {
		if response, err := handlers.GetAssets(ctx.Request.URL.Query()); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, response)
		}
	})
	server.GET("/token", func(ctx *gin.Context) {
		if response, err := utils.GetToken(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"token": response,
			})
		}
	})
	server.GET("/assetTicker", func(ctx *gin.Context) {
		if response, err := handlers.GetAssetTicker(ctx.Request.URL.Query()); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, response)
		}
	})
	server.Run(":8080")
}
