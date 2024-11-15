package src

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BEFORE = "2024-05-25"
	AFTER  = "2024-05-30"
)

func Test_SubDaya(t *testing.T) {
	sub := GetSubDay(AFTER, BEFORE)

	assert.Equal(t, sub, 5)
}

func Test_SubHour(t *testing.T) {
	sub := GetSubHour(AFTER, BEFORE)

	assert.Equal(t, sub, 120)
}

func Test_SubSeconds(t *testing.T) {
	sub := GetSubSeconds(AFTER, BEFORE)

	assert.Equal(t, sub, 432000)
}

func Test_SubTodayUseStandard(t *testing.T) {

	tt := ParsingDate(GetToday())
	jh_birth, dg_birth := "10-01", "08-04"
	_jh := fmt.Sprintf("%d-%s", tt.Year(), jh_birth)
	_dg := fmt.Sprintf("%d-%s", tt.Year(), dg_birth)

	if _jh > _dg {
		fmt.Println("true")
	}

	if _dg > _jh {
		fmt.Println("true")
	}

	jh_sub := GetSubTodayUseStandard(_jh)
	dg_sub := GetSubTodayUseStandard(_dg)

	fmt.Println("동규 생일 남은 날 : ", dg_sub)
	fmt.Println("정혜 생일 남은 날 : ", jh_sub)

}
