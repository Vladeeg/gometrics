package gauge

import (
	"github.com/Vladeeg/gometrics/internal/repository"
)

var memory = repository.NewMemStorage[float64]()

func SetGauge(key string, value float64) {
	memory.Set(key, value)
}
