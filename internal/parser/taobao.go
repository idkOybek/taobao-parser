package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/idkOybek/taobao-parser/internal/models"
	"github.com/idkOybek/taobao-parser/pkg/httpclient"
)

// TaobaoParser реализует парсер для сайта Taobao
type TaobaoParser struct {
	client *httpclient.Client
}

// NewTaobaoParser создает новый экземпляр парсера Taobao
func NewTaobaoParser() *TaobaoParser {
	return &TaobaoParser{
		client: httpclient.NewClient(),
	}
}

// ParseProduct парсит страницу товара на Taobao
func (p *TaobaoParser) ParseProduct(url string) (*models.Product, error) {
	resp, err := p.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении страницы: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при парсинге HTML: %w", err)
	}

	product := &models.Product{}

	// Парсим заголовок
	product.Title = doc.Find(".tb-main-title").Text()

	// Парсим описание
	product.Description = doc.Find(".tb-item-info-l").Text()

	// Парсим цену
	priceStr := doc.Find(".tb-rmb-num").First().Text()
	price, err := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)
	if err == nil {
		product.Price = price
	}

	// Парсим URL изображения
	product.ImageURL, _ = doc.Find("#J_ImgBooth").Attr("src")

	return product, nil
}
