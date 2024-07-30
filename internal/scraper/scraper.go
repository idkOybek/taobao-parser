package scraper

import (
	"github.com/idkOybek/taobao-parser/internal/models"
)

// Scraper определяет интерфейс для скраперов
type Scraper interface {
	ScrapeProduct(url string) (*models.Product, error)
	ScrapeCategory(url string, limit int) ([]*models.Product, error)
}
