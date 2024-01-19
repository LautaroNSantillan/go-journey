package fileuploadbot

import (
	"fmt"
	"os"

	"github.com/LautaroNSantillan/my-go-journey/slack-bot/file-upload-bot/pkg/fileuploadbot/definitions"
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"github.com/slack-io/slacker"
)

func InitializeBot() *slacker.Slacker {
	godotenv.Load()
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")

	return slacker.NewClient(botToken, appToken)
}

func ApiChanFilesSetup() (*slack.Client, []string, []string) {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("SLACK_CHANNEL_ID")}
	fileArr := []string{"../bono.jpg"}
	return api, channelArr, fileArr
}

func UploadFiles() {
	fmt.Println("Uploading files...")
	api, channelArr, fileArr := ApiChanFilesSetup()
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
			ctx.Response().Reply("🤖Here are your files!🤖")
			UploadFiles()
		},
	})

}
