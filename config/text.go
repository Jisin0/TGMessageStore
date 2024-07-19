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
<i>👋 Hi <b>{mention}</b> I'm a bot that can create <b>links</b> to <b>store</b> your <b>files and messages</b> in a single click 🔗
<b>Learn</b> how to use me 👉 /help to easily <b>share</b> content with anyone 👣 </i>
`,
	"ABOUT": `
○ <b>Language</b>: <a href='https://go.dev'>GO</a>
○ <b>Library</b>: <a href='https://github.com/PaulSonOfLars/gotgbot'>GoTgbot</a>
○ <b>Support</b>: <a href='https://t.me/FractalProjects'>@Fractal</a>
	`,

	"HELP": `
<i>👋 Hey {mention} I'm a bot that can create <b>permanent</b> links to a single or a <b>batch</b> of messages.</i>
<i><b>Add</b> me to your <b>channel</b> where you'll posting all the content before using me 😉</i>

<i>Here's a list of my available commands 👉</i>

/start : Start the bot.
/batch : Create a new message batch.
/genlink : Create link for a single post.
/about : Get some data about the bot.
/help  : Display this help message.
/privacy: Learn how this bot uses your data.
`,

	"PRIVACY": `<i>This bot does not connect to any datbase and hence <b>does not store any user data</b> in any form.</i>`,
}

// Message that is sent when an unrecognized command is sent.
var CommandNotFound = "<i>😐 I don't recognize that command !\nCheck /help to see how to use me.</i>"

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
<code>/batch start_post_link end_post_link</code>
<b>Example</b>
<code>/batch https://t.me/c/123456789/69 https://t.me/c/123456789/100</code>
</blockquote>`

	// Unable to access source channel
	BatchUnknownChat = "<i>🫤 I <b>couldn't access</b> that channel please make sure I am an <b>admin</b> there or <b>send a new message</b> if the channel is inactive !</i>"

	// Batch link was successfully generated.
	BatchSuccess = "<i>🎉 Here is your link :</i>\n<code>{link}</code>\n<a href='{link}'>Tap To Open</a>"

	// Batch exceeds size limit.
	BatchTooLarge = "<i>🫣 You can't make a batch that big my limit is {limit} !</i>"
)

// Genlink command texts.
// Error and success messages are same as batch.
var (
	GenlinkBadUsage = `<i>🤧 Command Usage was <b>Incorrect</b> !</i>
<blockquote expandable>
<b>Usage</b>
Add the bot to your channel and forward the post and use this command as a reply or copy the link of the post from the channel;
<b>Format</b>
<code>/genlink post_link</code>
<b>Example</b>
<code>/genlink https://t.me/c/123456789/69</code>
</blockquote>`
)

// Miscellaneous.
var (
	// malformed start link
	InvalidLink = "<i>I'm sorry there's something wrong with this link 😕</i>"
	// fetching batch messages
	StartGetBatch = "<i><b>Fetching your content...</b></i>"
)

// GetCommand returns the content for a command.
func GetCommand(command string) (string, [][]gotgbot.InlineKeyboardButton) {
	command = strings.ToUpper(command)

	text, ok := Commands[command]
	if !ok {
		text = CommandNotFound // default msg if not found
	}

	return text, Buttons[command]
}

// GetCommandText returns only text for a command.
