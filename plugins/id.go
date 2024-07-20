// (c) Jisin0
//
// id command handlers.

package plugins

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// ID handles the id command.
func ID(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage
	sender := ctx.EffectiveSender

	var output strings.Builder

	output.WriteString(fmt.Sprintf("<b>Sender ID</b> : <code>%d</code>\n", sender.Id()))

	if reply := update.ReplyToMessage; reply != nil {
		if forward := reply.ForwardOrigin; forward != nil {
			merged := forward.MergeMessageOrigin()

			if merged.Chat != nil {
				output.WriteString(fmt.Sprintf("<b>Forwarded From</b> : <code>%d</code>\n", merged.Chat.Id))
			}

			if merged.SenderChat != nil {
				output.WriteString(fmt.Sprintf("<b>Forwarded Group</b> : <code>%d</code>\n", merged.SenderChat.Id))
			}

			if merged.SenderUser != nil {
				output.WriteString(fmt.Sprintf("<b>Forwarded User</b> : <code>%d</code>\n", merged.SenderUser.Id))
			}

			if merged.SenderUserName != "" {
				output.WriteString(fmt.Sprintf("<b>Forwarded Username</b> : <code>%s</code>\n", merged.SenderUserName))
			}
		}
	}

	update.Reply(bot, output.String(), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})

	return ext.EndGroups
}
