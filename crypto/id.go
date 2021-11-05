package crypto

import (
	"math/rand"
	"time"
)

func GenerateId() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 3000
	return rand.Intn(max-min) + min
}
