package httpclient

import (
	"net/http"
	"time"
)

// Client представляет собой настраиваемый HTTP-клиент
type Client struct {
	*http.Client
}

// NewClient создает новый экземпляр HTTP-клиента с настроенными параметрами
func NewClient() *Client {
	return &Client{
		Client: &http.Client{
			Timeout: 30 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        100,
				MaxIdleConnsPerHost: 100,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

// Get выполняет GET-запрос к указанному URL
func (c *Client) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Устанавливаем заголовки для имитации браузера
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")

	return c.Do(req)
}
