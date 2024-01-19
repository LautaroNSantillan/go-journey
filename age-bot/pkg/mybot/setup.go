package mybot

import (
	"fmt"
	"os"

	"github.com/LautaroNSantillan/my-go-journey/tree/slack-bot/pkg/mybot/handlers"
	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func InitializeBot() *slacker.Slacker {
	godotenv.Load()
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	return slacker.NewClient(botToken, appToken)
}

func RegisterCommands(bot *slacker.Slacker) {

	bot.Command("Hello", &slacker.CommandDefinition{
		Description: "Greet the bot",
		Handler:     handlers.Greet,
	})

	bot.Command("My year of birth is <year>", &slacker.CommandDefinition{
		Description: "BEEP BOOP Age calculator with only the year BEEP BOOP\nExample: My year of birth is 1969",
		Handler:     handlers.YearAge,
	})

	bot.Command("My date of birth is <dd> <mm> <yyyy> ", &slacker.CommandDefinition{
		Description: "BEEP BOOP Actual age calculator, please enter your date of birth with the format dd/mm/yyyy BEEP BOOP",
		Handler:     handlers.ProperAge,
	})
}

func PrintCommandEvents(analiticsChannel <-chan *slacker.CommandEvent) {
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
