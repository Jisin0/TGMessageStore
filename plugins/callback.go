// (c) Jisin0
//
// Callback handlers.

package plugins

import (
	"fmt"
	"strings"

	"github.com/Jisin0/TGMessageStore/config"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// CbCommand handles callback from command buttons.
func CbCommand(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.CallbackQuery

	split := strings.SplitN(update.Data, "_", 2)
	if len(split) < 2 {
		update.Answer(bot, &gotgbot.AnswerCallbackQueryOpts{Text: "Bad Callback Data !", ShowAlert: true})
		return nil
	}

	var (
		cmd = strings.ToUpper(split[1])
	)

	text, ok := config.Commands[cmd]
	if !ok {
		text = config.CommandNotFound
	}

	_, _, err := update.Message.EditText(bot, text, &gotgbot.EditMessageTextOpts{ParseMode: gotgbot.ParseModeHTML, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: config.Buttons[cmd]}, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}})
	if err != nil {
		fmt.Println(err)
	}

	return ext.EndGroups
}
