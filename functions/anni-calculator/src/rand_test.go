package src

import (
	"fmt"
	"testing"
)

func Test_getRnad(t *testing.T) {
	r1 := GetRand(0, 9)
	r2 := GetRand(1, 10)
	r3 := GetRand(1, 10)
	r4 := GetRand(1, 10)
	r5 := GetRand(1, 10)

	fmt.Println(r1, r2, r3, r4, r5)
}
