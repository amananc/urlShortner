package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
	"urlshortner/constants"
	"urlshortner/internal/models"
)

func ExtractHashValue(url string) string {
	hashValue := ""

	if strings.HasPrefix(url, constants.BaseUrl) {
		hashValue = strings.TrimPrefix(url, constants.BaseUrl)
	} else if strings.HasPrefix(url, constants.ShortBaseUrl) {
		hashValue = strings.TrimPrefix(url, constants.ShortBaseUrl)
	}

	return hashValue
}

func IsShortUrlValid(url string) bool {
	if len(url) <= len(constants.BaseUrl) {
		return false
	}
	return strings.HasPrefix(url, constants.BaseUrl) || strings.HasPrefix(url, constants.ShortBaseUrl)
}

// todo : will move to a different file later
func CreateSuccessResponse(c *gin.Context, result models.URL) {
	response := gin.H{
		"message": constants.BaseUrl + result.Hash,
		"url":     result.OriginalUrl,
	}
	c.JSON(http.StatusOK, response)
}

func CreateErrorResponse(c *gin.Context, req models.RequestDto) {
	response := gin.H{
		"message": "Something went wrong",
		"url":     req.URL,
	}
	c.JSON(http.StatusInternalServerError, response)
}

func HandleBadRequest(c *gin.Context, message string, err error) {
	log.Printf("Bad request: %s - %v\n", message, err)
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   "Invalid shorten url",
		"details": err.Error(),
	})
}

func HandleNotFound(c *gin.Context, message string) {
	log.Printf("Not found: %s\n", message)
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Invalid shorten url",
	})
}

func CreateShortUrl(hash string, originalUrl string) models.URL {
	shortUrl := models.URL{
		Id:          uuid.New().String(),
		Hash:        hash,
		OriginalUrl: originalUrl,
		CreatedAt:   time.Now().String(),
	}

	return shortUrl
}

func CreateResponse(c *gin.Context, req models.RequestDto, result models.URL, err error) {
	if err != nil {
		CreateErrorResponse(c, req)
		return
	}
	CreateSuccessResponse(c, result)
}

func ParseRequest(c *gin.Context) (models.RequestDto, error) {
	var req models.RequestDto
	err := c.ShouldBindJSON(&req)
	if err != nil {
		HandleBadRequest(c, "Invalid request data", err)
	}

	return req, err
}
