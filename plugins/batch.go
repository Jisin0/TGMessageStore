// (c) Jisin0
//
// plugins/batch.go contains /batch command handlers and helpers.

package plugins

import (
	"errors"
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

// Batch handles the /batch command.
func Batch(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.Message
	user := ctx.EffectiveUser

	if !auth.CheckUser(user.Id) {
		update.Reply(bot, format.BasicFormat(config.BatchUnauthorized, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	args := strings.Fields(update.Text)
	if len(args) < 3 {
		update.Reply(bot, format.BasicFormat(config.BatchBadUsage, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	chatString, startID, err1 := parsePostLink(args[1])
	_, endID, err2 := parsePostLink(args[2])

	if err1 != nil || err2 != nil {
		update.Reply(bot, format.BasicFormat(config.BatchBadUsage, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	if startID > endID {
		update.Reply(bot, "Please enter the first post link before the last !", &gotgbot.SendMessageOpts{})
		return nil
	}

	if endID-startID > config.BatchSizeLimit {
		update.Reply(bot, format.BasicFormat(config.BatchTooLarge, user, map[string]any{"limit": config.BatchSizeLimit}), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})
		return nil
	}

	chatID, err := strconv.ParseInt(chatString, 10, 64)
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

	link := fmt.Sprintf("https://t.me/%s?start=%s", bot.Username, url.EncodeData(chatID, startID, endID))

	update.Reply(bot, format.BasicFormat(config.BatchSuccess, user, map[string]any{"link": link}), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML})

	return ext.EndGroups
}

// parsePostLink returns the username/id of the chat and the messageid from a link.
func parsePostLink(s string) (chatID string, messageID int64, err error) {
	split := strings.Split(s, "/")
	if len(split) < 3 {
		return chatID, messageID, errors.New("not enought url paths")
	}

	messageID, err = strconv.ParseInt(split[len(split)-1], 10, 64)

	chatID = split[len(split)-2]

	return chatID, messageID, err
}

// fixChatID adds a -100 to the start of a chatID assuming it's from a channel/supergroup.
func fixChatID(n int64) int64 {
	s := fmt.Sprint(n)
	if strings.HasPrefix(s, "-100") {
		return n
	}

	s = "-100" + s
	n, _ = strconv.ParseInt(s, 10, 64)

	return n
}
