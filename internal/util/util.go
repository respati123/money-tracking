package util

import (
	"math/rand"
	"time"
)

func GenerateNumber(digit int) int {
	if digit <= 0 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())
	min := 1
	for i := 1; i < digit; i++ {
		min *= 10
	}
	max := min*10 - 1

	return rand.Intn(max-min+1) + min

}
