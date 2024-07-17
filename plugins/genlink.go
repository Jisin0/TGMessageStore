// (c) Jisin0
//
// genlink command handlers and helpers.

package plugins

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Jisin0/TGMessageStore/config"
	"github.com/Jisin0/TGMessageStore/utils/auth"
	"github.com/Jisin0/TGMessageStore/utils/format"
	"github.com/Jisin0/TGMessageStore/utils/helpers"
	"github.com/Jisin0/TGMessageStore/utils/url"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// GenLink handles the /genlink command.
func GenLink(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage
	user := ctx.EffectiveUser

	if !auth.CheckUser(user.Id) {
		update.Reply(bot, format.BasicFormat(config.BatchUnauthorized, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	var (
		messageID int64
		chatID    int64
		err       error
	)

	if update.ReplyToMessage != nil && update.ReplyToMessage.ForwardOrigin != nil && update.ReplyToMessage.ForwardOrigin.MergeMessageOrigin().Chat != nil {
		messageID = update.ReplyToMessage.ForwardOrigin.MergeMessageOrigin().MessageId
		chatID = update.ReplyToMessage.ForwardOrigin.MergeMessageOrigin().Chat.Id
	} else {
		split := strings.Fields(update.Text)
		if len(split) < 2 {
			update.Reply(bot, format.BasicFormat(config.GenlinkBadUsage, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
			return nil
		}

		var chatString string

		chatString, messageID, err = parsePostLink(split[1])
		if err != nil {
			update.Reply(bot, format.BasicFormat(config.GenlinkBadUsage, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
			return nil
		}

		// Try to access the chat
		chatID, err = strconv.ParseInt(chatString, 10, 64)
		if err != nil {
			chatID, err = helpers.IDFromUsername(bot, chatString)
			if err != nil {
				update.Reply(bot, config.BatchUnknownChat, &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
				return nil
			}
		} else {
			chatID = fixChatID(chatID)

			_, err := bot.GetChat(chatID, &gotgbot.GetChatOpts{})
			if err != nil {
				update.Reply(bot, config.BatchUnknownChat, &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
				return nil
			}
		}
	}

	link := fmt.Sprintf("https://t.me/%s?start=%s", bot.Username, url.EncodeData(chatID, messageID, messageID))

	update.Reply(bot, format.BasicFormat(config.BatchSuccess, user, map[string]string{"link": link}), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
	return ext.EndGroups
}
