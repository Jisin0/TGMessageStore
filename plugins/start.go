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

	if len(config.FsubChannels) > 0 {
		var toJoin []*gotgbot.ChatFullInfo

		for _, c := range config.FsubChannels {
			if !isMember(bot, c, user.Id) {
				chat, err := bot.GetChat(c, &gotgbot.GetChatOpts{})
				if err != nil {
					continue
				}

				toJoin = append(toJoin, chat)
			}
		}

		if len(toJoin) > 0 {
			var buttons [][]gotgbot.InlineKeyboardButton

			switch len(toJoin) {
			case 1:
				buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: "ᴊᴏɪɴ ᴍʏ ᴄʜᴀɴɴᴇʟ", Url: toJoin[0].InviteLink}})
			case 2:
				buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: "ᴊᴏɪɴ ғɪʀsᴛ ᴄʜᴀɴɴᴇʟ", Url: toJoin[0].InviteLink}}, []gotgbot.InlineKeyboardButton{{Text: "ᴊᴏɪɴ sᴇᴄᴏɴᴅ ᴄʜᴀɴɴᴇʟ", Url: toJoin[1].InviteLink}})
			default:
				for i, c := range toJoin {
					buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: fmt.Sprintf("ᴊᴏɪɴ ᴄʜᴀɴɴᴇʟ %d", i+1), Url: c.InviteLink}})
				}
			}

			buttons = append(buttons, []gotgbot.InlineKeyboardButton{{Text: "ʀᴇᴛʀʏ 🔃", Url: fmt.Sprintf("https://t.me/%s?start=%s", bot.Username, split[1])}})

			update.Reply(bot, format.BasicFormat(config.FsubMessage, user), &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}})

			return nil
		}
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

	if endID-startID > config.BatchSizeLimit {
		statMessage.EditText(bot, format.BasicFormat(config.BatchTooLarge, inputMessage.From, map[string]any{"limit": config.BatchSizeLimit}), &gotgbot.EditMessageTextOpts{ParseMode: gotgbot.ParseModeHTML})
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

// isMember checks if a usr is a member of a chat.
func isMember(bot *gotgbot.Bot, chatID, userID int64) bool {
	m, err := bot.GetChatMember(chatID, userID, &gotgbot.GetChatMemberOpts{})
	if err != nil && strings.Contains(err.Error(), "user not found") {
		return false
	} else if err != nil {
		fmt.Printf("ismember: %v\n", err)
		return true
	}

	member := m.MergeChatMember()

	switch status := member.Status; status {
	case "left", "kicked", "banned":
		return false
	case "restricted":
		return member.IsMember
	}

	return true
}
