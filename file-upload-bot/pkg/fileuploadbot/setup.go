package fileuploadbot

import (
	"fmt"
	"os"

	"github.com/LautaroNSantillan/my-go-journey/slack-bot/file-upload-bot/pkg/fileuploadbot/definitions"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
)

var (
	api        *slack.Client
	channelArr []string
	fileArr    []string
	appToken   string
	botToken   string
)

func InitializeBot() *slacker.Slacker {
	godotenv.Load()
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	return slacker.NewClient(botToken, appToken)
}

func ApiChanFilesSetup() {
	api = slack.New(botToken)
	channelArr = []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr = []string{"./pkg/fileuploadbot/bono.jpg"}
}

func UploadFiles() {
	fmt.Println("Uploading files...")
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		fmt.Println(file)
		if err != nil {
			fmt.Printf("Error uploading file %s: %s\n", fileArr[i], err)
			return
		}
		fmt.Printf("Name: %s\n, URL: %s\n", file.Name, file.URL)
	}
}

func RegisterCommands(bot *slacker.Slacker) {

	bot.AddCommand(definitions.GreetDefinition)

	bot.AddCommand(&slacker.CommandDefinition{
		Command: "Upload",
		Handler: func(ctx *slacker.CommandContext) {
			UploadFiles()
			ctx.Response().Reply("🤖Here are your files!🤖")
		},
	})

}
