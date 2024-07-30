package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/idkOybek/taobao-parser/internal/config"
	"github.com/idkOybek/taobao-parser/internal/scraper"
)

func main() {
	cfg := config.NewConfig()
	scraper := scraper.NewCollyScraper()

	// Пример использования: парсинг отдельного товара с повторными попытками
	productURL := cfg.BaseURL + "/item/123456789.htm"
	product, err := scraper.ScrapeProductWithRetry(productURL, 3, 5*time.Second)
	if err != nil {
		log.Fatalf("Ошибка при парсинге товара: %v", err)
	}

	// Вывод результата в JSON
	jsonProduct, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при маршалинге JSON: %v", err)
	}
	fmt.Println(string(jsonProduct))

	// Пример использования: парсинг категории товаров с повторными попытками
	categoryURL := cfg.BaseURL + "/category/123.htm"
	products, err := scraper.ScrapeCategoryWithRetry(categoryURL, 10, 3, 5*time.Second)
	if err != nil {
		log.Fatalf("Ошибка при парсинге категории: %v", err)
	}

	// Вывод результатов в JSON
	jsonProducts, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при маршалинге JSON: %v", err)
	}
	fmt.Println(string(jsonProducts))
}
