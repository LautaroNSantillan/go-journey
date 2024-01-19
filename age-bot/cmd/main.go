package main

import (
	"context"
	"log"

	"github.com/LautaroNSantillan/my-go-journey/slack-bot/age-bot/pkg/mybot"
)

func main() {

	bot := mybot.InitializeBot()

	mybot.RegisterCommands(bot)

	go mybot.PrintCommandEvents(bot.CommandEvents())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
