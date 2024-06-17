package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	TIME_LAYOUT = "2006-01-02"

	START   = os.Getenv("START_DAY")
	BIRTH   = os.Getenv("BIRTH_DAY")
	WEBHOOK = "https://hooks.slack.com/services/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"
)

func HandleRequest(ctx context.Context, e json.RawMessage) (*int, error) {
	start := PieceOfYMD(START)
	today := GetToday()

	fmt.Println("Start : ", start)
	fmt.Println("Today : ", today)

	// 오늘은 몇일 만났는지?
	todayFromBeginDay := TimeSub(start, today)

	// 99일전 알림
	is99 := false
	if (todayFromBeginDay+1)%100 == 0 {
		is99 = true
	}

	// 100일전 알림
	is100 := false
	if todayFromBeginDay%100 == 0 {
		is100 = true
	}

	// 생일까지 몇일 남았는가?
	leastForBirth := TimeSubBirthday(BIRTH, today)

	fmt.Println("몇일 만났나요 ? ", todayFromBeginDay+1)
	fmt.Println("%99", is99)
	fmt.Println("%100", is100)
	fmt.Println("생일까지 남은 날짜 ? ", leastForBirth)

	sendToSlackTemplate(strconv.Itoa(todayFromBeginDay+1), is99, is100, strconv.Itoa(leastForBirth))
	return nil, nil
}

func sendToSlackTemplate(afterBeginDay string, is99 bool, is100 bool, afterBirthday string) {

	attac := slack.Attachment{}
	attac.AddField(slack.Field{Title: "--랑 몇일 만났니?", Value: afterBeginDay})

	if is99 {
		attac.AddField(slack.Field{Title: "*00일 하루 전 인가요?", Value: "Yes"})
	}

	if is100 {
		attac.AddField(slack.Field{Title: "오늘은 00일 입니다", Value: "Yes"})
	}

	attac.AddField(slack.Field{Title: "-- 남은 생일 날", Value: afterBirthday})

	now := GetToday()
	p := slack.Payload{
		Text:        now,
		Channel:     "#chatbot-test",
		Attachments: []slack.Attachment{attac},
	}

	err := slack.Send(WEBHOOK, "", p)
	if err != nil {
		panic(err)
	}
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
