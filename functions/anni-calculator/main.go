package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ashwanthkumar/slack-go-webhook"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zkfmapf123/anni-calculator/src"
)

var (
	FIRST_MEET  = "2024-05-25" // 정혜랑 처음만난 날
	JH_BIRTHDAY = "10-01"      // 정혜 생일
	DG_BIRTHDAY = "08-04"      // 동규 생일
)

var (
	WEBHOOK     = "https://hooks.slack.com/services/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"
	CHANNEL     = "#chatbot-test"
	ImojiRandom = []string{
		":c1:", ":c2:", ":c3:", ":c4:", ":c5:",
		":c6:", ":c7:", ":c8:", ":c9:", ":c10:",
	}
)

type ILoveYouHJH struct {
	today       string // 오늘 시간
	meetDay     int    // 만난 날 (day)
	meetHour    int    // 만난 시간 (hour)
	meetSeconds int    // 만난 초 (seconds)

	remainDayJHBirth int // 정혜 생일 남은 날
	remainDayDGBirth int // 동규 생일 남은 날
}

func HandleRequest(ctx context.Context, e json.RawMessage) (*int, error) {

	iloveHjh := ILoveYouHJH{}

	today := src.GetToday()

	iloveHjh.today = today
	iloveHjh.meetDay = src.GetSubDay(today, FIRST_MEET)
	iloveHjh.meetHour = src.GetSubHour(today, FIRST_MEET)
	iloveHjh.meetSeconds = src.GetSubSeconds(today, FIRST_MEET)

	t := src.ParsingDate(src.GetToday())
	_jh_birthday := fmt.Sprintf("%d-%s", t.Year(), JH_BIRTHDAY)
	_dg_birthday := fmt.Sprintf("%d-%s", t.Year(), DG_BIRTHDAY)

	iloveHjh.remainDayJHBirth = src.GetSubTodayUseStandard(_jh_birthday)
	iloveHjh.remainDayDGBirth = src.GetSubTodayUseStandard(_dg_birthday)

	sendToSlackTemplate(iloveHjh)

	return nil, nil
}

// ////////////////////////////////////////////////////// Send Slack ////////////////////////////////////////////////////////
func sendToSlackTemplate(iloveyou ILoveYouHJH) {

	attac := slack.Attachment{}
	attac.AddField(slack.Field{Title: "정혜랑 몇일 만났니?", Value: fmt.Sprintf("%d 일", iloveyou.meetDay)})
	attac.AddField(slack.Field{Title: "정혜 생일 남은 날", Value: fmt.Sprintf("%d 일", iloveyou.remainDayJHBirth), Short: true})
	attac.AddField(slack.Field{Title: "동규 생일 남은 날", Value: fmt.Sprintf("%d 일", iloveyou.remainDayDGBirth), Short: true})

	attac.AddField(slack.Field{Title: "", Value: ImojiRandom[src.GetRand(0, 9)], Short: true})
	// now := GetToday()
	p := slack.Payload{
		Text:        src.GetToday(),
		Channel:     "#chatbot-test",
		Attachments: []slack.Attachment{attac},
	}

	err := slack.Send(WEBHOOK, "", p)
	if len(err) > 0 {
		log.Fatalln(err)
	}
}

func main() {
	lambda.Start(HandleRequest)
}
