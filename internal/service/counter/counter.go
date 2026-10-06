package counter

import (
	"github.com/Vladeeg/gometrics/internal/repository"
)

var memory = repository.NewMemStorage[int64]()

func UpdateCounter(key string, value int64) {
	prevCount := memory.Get(key)

	memory.Set(key, prevCount + value)
}
