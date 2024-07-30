package utils

import (
	"fmt"
	"time"
)

// Retry выполняет функцию с заданным количеством попыток и интервалом между ними
func Retry(attempts int, sleep time.Duration, f func() error) error {
	var err error
	for i := 0; ; i++ {
		err = f()
		if err == nil {
			return nil
		}

		if i >= attempts-1 {
			break
		}

		time.Sleep(sleep)
		fmt.Printf("Повторная попытка %d после ошибки: %v\n", i+1, err)
	}
	return fmt.Errorf("после %d попыток: %w", attempts, err)
}
