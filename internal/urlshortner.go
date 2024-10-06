package internal

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type shortenRequestDto struct {
	URL string `json:"url"`
}

type redirectDto struct {
	URL string `json:"url"`
}

var shortenUrlMap = make(map[string]string)

const BASE_URL = "www.baseUrl.com/"
const SHORT_BASE_URL = "baseUrl.com/"

func GenerateShortenUrl(context *gin.Context) {
	fmt.Println("Inside generateShortenUrl")

	var request shortenRequestDto
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H {
			"error": err.Error(),
		})
		return
	}

    fmt.Println("URL:", request.URL)
	hash_value := generateHashValue(request.URL)
	fmt.Println("hash value : " + hash_value + " for url " + request.URL)
	
	shortenUrlMap[hash_value] = request.URL

    context.JSON(http.StatusOK, gin.H{
        "message": BASE_URL + hash_value,
        "url":     request.URL,
    })
}

func RedirectToOriginalUrl(context *gin.Context) {
	fmt.Println("inside redirectToOriginalUrl")

	var request redirectDto
	
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("requested url : " + request.URL)

	if !validateRequestUrl(request.URL) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Requested url not valid"})
		return
	}

	hash_value := extractHashValue(request.URL)
	original_url := shortenUrlMap[hash_value]

	fmt.Println("original_url : " + original_url)
	createSuccessResponse(context, original_url)
}

func createSuccessResponse(context *gin.Context, url string) {
	context.Status(http.StatusOK)
	context.JSON(http.StatusOK, gin.H{"original url" : url})
}

func validateRequestUrl(val string) bool {
	if len(val) <= len(BASE_URL) {
		return false
	}

	if !strings.HasPrefix(val, BASE_URL) && !strings.HasPrefix(val, SHORT_BASE_URL) {
		return false
	}

	return true;
}

func extractHashValue(url string) string {
	hash_value := ""

	if strings.HasPrefix(url, BASE_URL) {
		hash_value = strings.TrimPrefix(url, BASE_URL)
	} else if strings.HasPrefix(url, SHORT_BASE_URL) {
		hash_value = strings.TrimPrefix(url, SHORT_BASE_URL)
	}

	fmt.Println("hash_value : " + hash_value)
	return hash_value
}

func GetAllMapping(context *gin.Context) {
	context.Status(http.StatusOK)
	context.JSON(http.StatusOK, shortenUrlMap)
}