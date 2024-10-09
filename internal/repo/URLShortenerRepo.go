package repo

import "urlshortner/internal/models"

type URLShortenerRepo interface {
	Save(request models.URL) (models.URL, error)
	GetAll() ([]models.URL, error)

	FindByHash(hash string) (models.URL, error)
	FindById(id string) (models.URL, error)

	DeleteByHash(hash string) (models.URL, error)
	Delete(id string) (models.URL, error)
}
