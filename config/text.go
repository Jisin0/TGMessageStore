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
	"START": `
Hey there partner !
`,
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

// Batch command texts.
var (
	// Unauthorized use of /batch
	BatchUnauthorized = "<i>😐 Sorry dude <b>only</b> an <b>admin</b> can do that !</i>"
	// Bad/Incorrect isage of /batch
	BatchBadUsage = `<i>🤧 Command Usage was <b>Incorrect</b> !</i>
<blockquote expandable>
<b>Usage</b>
Add the bot to your channel and copy the link of the first and last post(including) from the channel;
<b>Format</b>
<code>/batch <start_post_link> <end_post_link>
<b>Example</b>
<code>/batch https://t.me/c/123456789/69 https://t.me/c/123456789/100
</blockquote>`

	// Unable to access source channel
	BatchUnknownChat = "<i>🫤 I <b>couldn't access</b> that channel please make sure I am an <b>admin</b> there or <b>send a new message</b> if the channel is inactive !</i>"

	// Batch link was successfully generated.
	BatchSuccess = "<i>🎉 Here is your link :</i>\n<code>{link}</code>\n<a href='{link}'>Tap To Open</a>"
)
