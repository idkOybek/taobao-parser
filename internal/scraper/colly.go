package scraper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/idkOybek/taobao-parser/internal/models"
	"github.com/idkOybek/taobao-parser/pkg/utils"
)

// CollyScraper реализует скрапер на основе библиотеки Colly
type CollyScraper struct {
	collector *colly.Collector
}

// NewCollyScraper создает новый экземпляр скрапера Colly
func NewCollyScraper() *CollyScraper {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"),
	)
	return &CollyScraper{collector: c}
}

// ScrapeProduct выполняет скрапинг отдельного товара
func (s *CollyScraper) ScrapeProduct(url string) (*models.Product, error) {
	product := &models.Product{}

	s.collector.OnHTML(".tb-main-title", func(e *colly.HTMLElement) {
		product.Title = strings.TrimSpace(e.Text)
	})

	s.collector.OnHTML(".tb-item-info-l", func(e *colly.HTMLElement) {
		product.Description = strings.TrimSpace(e.Text)
	})

	s.collector.OnHTML(".tb-rmb-num", func(e *colly.HTMLElement) {
		priceStr := strings.TrimSpace(e.Text)
		price, err := strconv.ParseFloat(priceStr, 64)
		if err == nil {
			product.Price = price
		}
	})

	s.collector.OnHTML("#J_ImgBooth", func(e *colly.HTMLElement) {
		product.ImageURL = e.Attr("src")
	})

	err := s.collector.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при посещении страницы: %w", err)
	}

	return product, nil
}

// ScrapeCategory выполняет скрапинг товаров из категории
func (s *CollyScraper) ScrapeCategory(url string, limit int) ([]*models.Product, error) {
	products := []*models.Product{}

	s.collector.OnHTML(".item", func(e *colly.HTMLElement) {
		if len(products) >= limit {
			return
		}

		product := &models.Product{}
		product.Title = e.ChildText(".title")
		product.ImageURL = e.ChildAttr("img", "src")
		priceStr := e.ChildText(".price")
		price, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
		if err == nil {
			product.Price = price
		}

		products = append(products, product)
	})

	err := s.collector.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при посещении страницы категории: %w", err)
	}

	return products, nil
}

// ScrapeProductWithRetry выполняет скрапинг товара с механизмом повторных попыток
func (s *CollyScraper) ScrapeProductWithRetry(url string, attempts int, sleep time.Duration) (*models.Product, error) {
	var product *models.Product
	err := utils.Retry(attempts, sleep, func() error {
		var err error
		product, err = s.ScrapeProduct(url)
		return err
	})
	return product, err
}

// ScrapeCategoryWithRetry выполняет скрапинг категории с механизмом повторных попыток
func (s *CollyScraper) ScrapeCategoryWithRetry(url string, limit int, attempts int, sleep time.Duration) ([]*models.Product, error) {
	var products []*models.Product
	err := utils.Retry(attempts, sleep, func() error {
		var err error
		products, err = s.ScrapeCategory(url, limit)
		return err
	})
	return products, err
}
