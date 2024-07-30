package parser

import (
	"github.com/idkOybek/taobao-parser/internal/models"
)

// Parser определяет интерфейс для парсеров
type Parser interface {
	ParseProduct(url string) (*models.Product, error)
}
