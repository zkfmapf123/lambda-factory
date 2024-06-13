package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	TIME_LAYOUT = "2006-01-02"

	START = os.Getenv("START_DAY")
	BIRTH = os.Getenv("BIRTH_DAY")
)

func HandleRequest(ctx context.Context, e json.RawMessage) (*int, error) {
	start := PieceOfYMD(START)
	today := GetToday()

	fmt.Println("Start : ", start)
	fmt.Println("Today : ", today)

	// 오늘은 몇일 사귀었는지?
	todayFromBeginDay := TimeSub(start, today)

	// 100일 전 알림
	is99 := false
	if (todayFromBeginDay+1)%100 == 0 {
		is99 = true
	}

	// x % 100
	is100 := false
	if todayFromBeginDay%100 == 0 {
		is100 = true
	}

	// 생일까지 몇일 남았는가?
	leastForBirth := TimeSubBirthday(BIRTH, today)

	fmt.Println("몇일 만났나요 ? ", todayFromBeginDay)
	fmt.Println("%99", is99)
	fmt.Println("%100", is100)
	fmt.Println("생일까지 남은 날짜 ? ", leastForBirth)

	return nil, nil
}

// yyyy-mm-dd
func PieceOfYMD(v string) string {
	str := strings.Split(v, "$")
	joinStr := strings.Join(str, "-")
	return joinStr
}

// yyyy-mm-dd
func GetToday() string {
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatalln(err)
	}

	now := time.Now().In(loc)
	return now.Format("2006-01-02")
}

func TimeSubBirthday(birthMMDD string, today string) int {
	curYear := time.Now().Year()
	fullDate := fmt.Sprintf("%d-%s", curYear, birthMMDD)

	return TimeSub(fullDate, today)
}

func TimeSub(start, end string) int {
	startDate, _ := time.Parse(TIME_LAYOUT, start)
	endDate, _ := time.Parse(TIME_LAYOUT, end)

	diff := endDate.Sub(startDate)
	return int(diff.Hours() / 24)
}

func main() {
	lambda.Start(HandleRequest)
}
