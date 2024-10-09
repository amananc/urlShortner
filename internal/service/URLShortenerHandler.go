package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"urlshortner/internal/models"
	"urlshortner/internal/repo"
	"urlshortner/pkg/utils"
)

type URLShortenerHandler struct {
	repo repo.URLShortenerRepo
}

func NewURLShortenerHandler(repo repo.URLShortenerRepo) *URLShortenerHandler {
	return &URLShortenerHandler{repo: repo}
}

func (h *URLShortenerHandler) CreateURL(c *gin.Context) {
	log.Println("inside createUrl method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	hashValue := generateHashValue(req.URL)
	result, err := h.repo.Save(utils.CreateShortUrl(hashValue, req.URL))
	utils.CreateResponse(c, req, result, err)
}

func (h *URLShortenerHandler) Redirect(c *gin.Context) {
	log.Println("inside redirect method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	if !utils.IsShortUrlValid(req.URL) {
		utils.HandleNotFound(c, "Requested URL is not valid")
		return
	}

	hashValue := utils.ExtractHashValue(req.URL)

	var result models.URL
	result, err = h.repo.FindByHash(hashValue)
	if err != nil {
		utils.HandleNotFound(c, "Requested URL is not found")
		return
	}
	c.Redirect(http.StatusFound, result.OriginalUrl)
}

func (h *URLShortenerHandler) GetAll(c *gin.Context) {
	log.Println("inside getAll method")

	result, err := h.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *URLShortenerHandler) FindByHash(c *gin.Context) {
	log.Println("inside findByHash method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	result, err := h.repo.FindByHash(req.Hash)
	utils.CreateResponse(c, req, result, err)
}

func (h *URLShortenerHandler) FindById(c *gin.Context) {
	log.Println("inside findById method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	result, err := h.repo.FindById(req.Id)
	utils.CreateResponse(c, req, result, err)
}

func (h *URLShortenerHandler) DeleteByHash(c *gin.Context) {
	log.Println("inside deleteByHash method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	result, err := h.repo.DeleteByHash(req.Hash)
	utils.CreateResponse(c, req, result, err)
}

func (h *URLShortenerHandler) Delete(c *gin.Context) {
	log.Println("inside deleteByHash method")

	req, err := utils.ParseRequest(c)
	if err != nil {
		return
	}

	result, err := h.repo.Delete(req.Id)
	utils.CreateResponse(c, req, result, err)
}
