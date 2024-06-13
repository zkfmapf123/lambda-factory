package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	birth = "01-10"
	start = "2024$05$26"
)

func Test_PieceOfYMD(t *testing.T) {
	birth := PieceOfYMD(birth)
	start := PieceOfYMD(start)

	assert.Equal(t, birth, "1989-01-10")
	assert.Equal(t, start, "2024-05-26")
}

func Test_TimeSub(t *testing.T) {
	today := GetToday()
	_start := PieceOfYMD(start)

	sub := TimeSub(_start, today)
	fmt.Println(sub)
}

func Test_SubBirthday(t *testing.T) {
	birthMMDD := "10-01"
	today := GetToday()

	sub := TimeSubBirthday(birthMMDD, today)
	fmt.Println(sub)
}
