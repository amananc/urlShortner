package repo

import (
	"gorm.io/gorm"
	"log"
	"urlshortner/internal/models"
)

type UrlShortenerRepo struct {
	db *gorm.DB
}

func NewURLShortenerRepo(db *gorm.DB) *UrlShortenerRepo {
	log.Println("Initializing URLShortenerRepo")
	return &UrlShortenerRepo{db: db}
}

func (repo *UrlShortenerRepo) Save(request models.URL) (models.URL, error) {
	log.Println("Saving shortener")
	result := repo.db.Save(request)
	return request, result.Error
}

func (repo *UrlShortenerRepo) GetAll() ([]models.URL, error) {
	log.Println("Getting all")

	var urls []models.URL
	result := repo.db.Find(&urls)
	return urls, result.Error
}

func (repo *UrlShortenerRepo) FindByHash(hash string) (models.URL, error) {
	log.Println("inside findByHash")

	var url models.URL
	result := repo.db.Where("hash = ?", hash).First(&url)
	return url, result.Error
}

func (repo *UrlShortenerRepo) FindById(id string) (models.URL, error) {
	log.Println("inside findById")

	var url models.URL
	result := repo.db.Where("id = ?", id).First(&url)
	return url, result.Error
}

func (repo *UrlShortenerRepo) DeleteByHash(hash string) (models.URL, error) {
	log.Println("Deleting shortener by hash")

	var url models.URL
	result := repo.db.Where("hash = ?", hash).Delete(&url)
	return url, result.Error
}

func (repo *UrlShortenerRepo) Delete(id string) (models.URL, error) {
	log.Println("Deleting shortener")

	var url models.URL
	result := repo.db.Where("id = ?", id).Delete(&url)
	return url, result.Error
}
