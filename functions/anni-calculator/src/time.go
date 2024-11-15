package src

import (
	"log"
	"time"
)

const (
	LAYOUT = "2006-01-02"
)

func ParsingDate(a string) time.Time {
	dateA, _ := time.Parse(LAYOUT, a)
	return dateA
}

func parsingDates(a, b string) (time.Time, time.Time) {
	dateA, _ := time.Parse(LAYOUT, a)
	dateB, _ := time.Parse(LAYOUT, b)

	return dateA, dateB
}

func GetSubDay(a, b string) int {

	tA, tB := parsingDates(a, b)

	duration := tA.Sub(tB)
	return int(duration.Hours() / 24)
}

func GetSubHour(a, b string) int {

	tA, tB := parsingDates(a, b)

	duration := tA.Sub(tB)
	return int(duration.Hours())
}

func GetSubSeconds(a, b string) int {
	tA, tB := parsingDates(a, b)

	duration := tA.Sub(tB)
	return int(duration.Seconds())
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

func GetSubTodayUseStandard(standard string) int {

	todayDate := ParsingDate(GetToday())
	standartDate := ParsingDate(standard)

	var newBirthYYYYMMDD time.Time

	// 생일이 이미 지남
	if todayDate.After(standartDate) {
		newBirthYYYYMMDD = time.Date(todayDate.AddDate(1, 0, 0).Year(), standartDate.Month(), standartDate.Day(), 0, 0, 0, 0, time.UTC)
	} else {
		newBirthYYYYMMDD = time.Date(todayDate.Year(), standartDate.Month(), standartDate.Day(), 0, 0, 0, 0, time.UTC)
	}

	return int(newBirthYYYYMMDD.Sub(todayDate).Hours()) / 24

}
