package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func main() {

	godotenv.Load()
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	api := slack.New(botToken)
	channelArr := []string{channelId}
	fileArr := []string{"bono.jpg"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s\n, URL: %s\n", file.Name, file.URL)
	}
}
