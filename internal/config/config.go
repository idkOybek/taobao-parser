package config

// Config содержит настройки приложения
type Config struct {
	BaseURL string
	Timeout int
}

// NewConfig создает новый экземпляр конфигурации с значениями по умолчанию
func NewConfig() *Config {
	return &Config{
		BaseURL: "https://world.taobao.com",
		Timeout: 30,
	}
}
