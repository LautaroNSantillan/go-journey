package handlers

import (
	"fmt"
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
	day := req.Param("day")
	month := req.Param("MM")
	year := req.Param("YYYY")

	dateStr := fmt.Sprintf("%s/%s/%s", day, month, year)
	parsedDate, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		res.Reply("Error: Invalid date format. Please use the format dd/mm/yyyy.")
	}

	currentTime := time.Now()
	ageDuration := currentTime.Sub(parsedDate)
	years := int(ageDuration.Hours() / 24 / 365.25)
	months := int(ageDuration.Hours() / 24 / 30.44)
	days := int(ageDuration.Hours() / 24)

	res.Reply(fmt.Sprintf("Your age is %d years, %d months, and %d days", years, months, days))
}
