package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func Greet(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
	res.Reply("Hello! I am a bot that calculates your age 🤖")
}

func YearAge(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
	year := req.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		res.Reply("Invalid year")
	}
	currentTime := time.Now().Year()
	age := currentTime - yob

	res.Reply(fmt.Sprintf("Your age is %d", age))
}

func ProperAge(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
	day := req.Param("dd")
	month := req.Param("mm")
	year := req.Param("yyyy")

	log.Printf("Day: %s, Month: %s, Year: %s", day, month, year)

	if strToInt(day) < 1 || strToInt(day) > 31 || strToInt(month) < 1 || strToInt(month) > 12 || strToInt(year) < 1900 {
		res.Reply("Error: Invalid date. Please enter a valid date.")
		return
	}

	dateStr := fmt.Sprintf("%s/%s/%s", day, month, year)
	fmt.Println("DATESTR : ", dateStr)
	parsedDate, err := time.Parse("2/1/2006", dateStr)
	fmt.Println("PARSED DATE ", parsedDate.String())
	if err != nil {
		res.Reply("Error: Invalid date format. Please use the format dd mm yyyy.")
	}

	years, months, days := diffInYearsMonthsDays(parsedDate, time.Now())

	res.Reply(fmt.Sprintf("Your age is %d years, %d months, and %d days", years, months, days))
}

func strToInt(str string) int {
	int, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	fmt.Println("strToInt: ", int)
	return int
}

func diffInYearsMonthsDays(start, end time.Time) (int, int, int) {
	years := end.Year() - start.Year()
	months := int(end.Month()) - int(start.Month())
	days := end.Day() - start.Day()

	if months < 0 || (months == 0 && days < 0) {
		years--
		months += 12
	}

	if days < 0 {

		lastDayOfPreviousMonth := start.AddDate(0, 1, 0).AddDate(0, 0, -start.Day())
		days = int(end.Sub(lastDayOfPreviousMonth).Hours() / 24)
	}

	return years, months, days
}
