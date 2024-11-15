package src

import (
	"math/rand"
)

func GetRand(min, max int) int {
	randNum := rand.Intn(max) + min
	return randNum
}
