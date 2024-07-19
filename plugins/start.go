// (c) Jisin0
//
// File contains start command handler and helpers.

package plugins

import (
	"fmt"
	"strings"

	"github.com/Jisin0/TGMessageStore/config"
	"github.com/Jisin0/TGMessageStore/utils/format"
	"github.com/Jisin0/TGMessageStore/utils/url"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage
	user := ctx.EffectiveUser

	split := strings.Fields(update.Text)
	if len(split) < 2 {
		text, buttons := config.GetCommand("START")
		update.Reply(bot, format.BasicFormat(text, user), &gotgbot.SendMessageOpts{ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}, ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	chatID, startID, endID, err := url.DecodeData(split[1])
	if err != nil {
		fmt.Println(err)
		update.Reply(bot, format.BasicFormat(config.InvalidLink, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	sendBatch(bot, update, chatID, startID, endID)

	return nil
}

// sendBatch sends a batch from the input data to the target.
func sendBatch(bot *gotgbot.Bot, inputMessage *gotgbot.Message, chatID, startID, endID int64) {
	statMessage, err := inputMessage.Reply(bot, format.BasicFormat(config.StartGetBatch, inputMessage.From), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
	if err != nil {
		fmt.Printf("sendBatch: failed to send stat: %v\n", err)
		return
	}

	for i := startID; i <= endID; i++ {
		_, err := bot.CopyMessage(inputMessage.Chat.Id, chatID, i, &gotgbot.CopyMessageOpts{ProtectContent: config.ProtectContent, DisableNotification: config.DisableNotification})
		if err != nil {
			switch {
			case strings.Contains(err.Error(), "chat not found"):
				statMessage.EditText(bot, format.BasicFormat(config.BatchUnknownChat, inputMessage.From), &gotgbot.EditMessageTextOpts{})
				return
			case strings.Contains(err.Error(), "message not found"):
				// ignore and continue
			case strings.Contains(err.Error(), "flood"):
				fmt.Println("cancelled batch due to flood")
				return
			default:
				fmt.Printf("sendBatch: unknown error: %v", err)
			}
		}
	}

	statMessage.Delete(bot, &gotgbot.DeleteMessageOpts{})
}
