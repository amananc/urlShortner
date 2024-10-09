package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortner/config"
	"urlshortner/internal/repo"
	"urlshortner/internal/service"
)

func InitializeApp() {
	fmt.Println("inside initializeApp")

	config.InitConfig()
	db := repo.InitializeDB()

	urlRepo := repo.NewURLShortenerRepo(db)
	handler := service.NewURLShortenerHandler(urlRepo)

	router := gin.Default()
	router.GET("/health", health)
	router.POST("/shorten", handler.CreateURL)
	router.GET("/redirect", handler.Redirect)
	router.POST("/all", handler.GetAll)

	// Testing apis
	router.DELETE("/delete", handler.Delete)
	router.DELETE("/deleteByHash", handler.DeleteByHash)
	router.GET("/findByHash", handler.FindByHash)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

func health(context *gin.Context) {
	fmt.Println("Status OK!")
	context.JSON(http.StatusOK, gin.H{"status": "OK"})
}
