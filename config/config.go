package config

import (
	"fmt"
	"net/http"
	"urlshortner/internal"
	"github.com/gin-gonic/gin"
)

func InitializeApp() {
	fmt.Println("inside initializeApp")

	router := gin.Default();

	router.GET("/test", test)
	router.POST("/shorten", internal.GenerateShortenUrl)
	router.POST("/redirect", internal.RedirectToOriginalUrl)
	router.POST("/all", internal.GetAllMapping)

	router.Run("localhost:8080")
}

func test(context *gin.Context) {
	fmt.Println("testing OK!")
	context.JSON(http.StatusOK, "testing OK!")
}