package main

import (
	"context"
	"log"

	"github.com/LautaroNSantillan/my-go-journey/slack-bot/file-upload-bot/pkg/fileuploadbot"
)

func main() {

	bot := fileuploadbot.InitializeBot()

	fileuploadbot.RegisterCommands(bot)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
