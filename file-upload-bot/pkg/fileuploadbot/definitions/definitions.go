package definitions

import (
	"github.com/slack-io/slacker"
)

var GreetDefinition = &slacker.CommandDefinition{
	Command: "Hello",
	Handler: func(ctx *slacker.CommandContext) {
		ctx.Response().Reply("Hello! I am a bot that uploads files to this channel 🤖")
	},
}
