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

	output.WriteString(fmt.Sprintf("Sender ID : <code>%d</code>", sender.Id()))

	if forward := update.ForwardOrigin; forward != nil {
		merged := forward.MergeMessageOrigin()

		if merged.Chat != nil {
			output.WriteString(fmt.Sprintf("Forwarded From : <code>%d</code>", merged.Chat.Id))
		}

		if merged.SenderChat != nil {
			output.WriteString(fmt.Sprintf("Forwarded Group : <code>%d</code>", merged.SenderChat.Id))
		}

		if merged.SenderUser != nil {
			output.WriteString(fmt.Sprintf("Forwarded User : <code>%d</code>", merged.SenderUser.Id))
		}

		if merged.SenderUserName != "" {
			output.WriteString(fmt.Sprintf("Forwarded Username : <code>%s</code>", merged.SenderUserName))
		}
	}

	update.Reply(bot, output.String(), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})

	return ext.EndGroups
}
