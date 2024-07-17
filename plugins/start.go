// (c) Jisin0
//
// File contains start command handler and helpers.

package plugins

import (
	"github.com/Jisin0/TGMessageStore/config"
	"github.com/Jisin0/TGMessageStore/utils/format"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage
	user := ctx.EffectiveUser

	update.Reply(bot, format.BasicFormat(config.Commands["START"], user), &gotgbot.SendMessageOpts{})
	return nil
}
