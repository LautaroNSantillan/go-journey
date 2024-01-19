package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load()
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	bot := slacker.NewClient(botToken, appToken)

	bot.Command("Hello", &slacker.CommandDefinition{
		Description: "Greet the bot",
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
			res.Reply("Hello! I am a bot that calculates your age 🤖")
		},
	})

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "BEEP BOOP Age calculator with only the year BEEP BOOP\nExample: My year of birth is 1969",
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
			year := req.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				res.Reply("Invalid year")
			}
			currentTime := time.Now().Year()
			age := currentTime - yob

			res.Reply(fmt.Sprintf("Your age is %d", age))
		},
	})

	bot.Command("My date of birth is <DD>/<MM>/<YYYY> ", &slacker.CommandDefinition{
		Description: "BEEP BOOP Actual age calculator, please enter your date of birth with the format dd/mm/yyyy BEEP BOOP",
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
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
		},
	})

	go printCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func printCommandEvents(analiticsChannel <-chan *slacker.CommandEvent) {
	for event := range analiticsChannel {
		fmt.Println("Command events :")
		fmt.Println()
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
