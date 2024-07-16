// (c) Jisin0
//
// plugins/batch.go contains /batch command handlers and helpers.

package plugins

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// Batch handles the /batch command.
func Batch(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.Message
	user := ctx.EffectiveUser

}
