// (c) Jisin0
//
// config/text.go contains constant texts used across different commands.

package config

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

// Standard command replies. Add a new entry to create new command no extra configs needed.
var Commands map[string]string = map[string]string{
	"ABOUT": `
○ <b>Language</b>: <a href='https://go.dev'>GO</a>
○ <b>Library</b>: <a href='https://github.com/PaulSonOfLars/gotgbot'>GoTgbot</a>
○ <b>Support</b>: <a href='https://t.me/FractalProjects'>@Fractal</a>
	`,

	"HELP": `
<i>👋 Hey {mention} I'm a bot that can create <b>permanent</b> links to a single or a <b>batch</b> of messages.</i>

<i>Here's a list of my available commands 👉</i>

/start : Start the bot.
/batch : Create a new message batch.
/genlink : Generate link for a single post.
/about : Get some data about the bot.
/help  : Display this help message.
/privacy: Leran how this bot uses your data.
`,

	"PRIVACY": `<i>This bot does not connect to any datbase and hence <b>does not store any user data</b> in any form.</i>`,
}

// Message that is sent when an unrecognized command is sent.
var CommandNotFound = "<i>😐 I don't recognize that command !\nCheck /help to see how to use me.</i>"

// GetCommand returns the content for a command.
func GetCommand(command string) (string, [][]gotgbot.InlineKeyboardButton) {
	command = strings.ToUpper(command)

	text, ok := Commands[command]
	if !ok {
		text = CommandNotFound // default msg if not found
	}

	return text, Buttons[command]
}
