package server

import (
	"github.com/SungminSo/qr-generator/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	qrCode := router.Group("/qr/")
	{
		qrCode.POST("/generate", pkg.SendQR)
		qrCode.GET("/verify/:qrToken", pkg.VerifyQR)
	}
}