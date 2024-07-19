// (c) Jisin0

package plugins

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/Jisin0/TGMessageStore/config"
	"github.com/Jisin0/TGMessageStore/utils/format"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
)

var Dispatcher *ext.Dispatcher = ext.NewDispatcher(&ext.DispatcherOpts{
	// If an error is returned by a handler, log it and continue going.
	Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
		fmt.Println("an error occurred while handling update:", err.Error())
		return ext.DispatcherActionNoop
	},
	MaxRoutines: ext.DefaultMaxRoutines,
})

const (
	commandHandlerGroup  = 2
	callbackHandlerGroup = 1
)

func init() {
	Dispatcher.AddHandlerToGroup(handlers.NewCommand("start", Start), commandHandlerGroup)
	Dispatcher.AddHandlerToGroup(handlers.NewCommand("batch", Batch), commandHandlerGroup)
	Dispatcher.AddHandlerToGroup(handlers.NewCommand("genlink", GenLink), commandHandlerGroup)
	Dispatcher.AddHandlerToGroup(handlers.NewMessage(allCommand, CommandHandler), commandHandlerGroup) // handles all other unhandled commands

	Dispatcher.AddHandlerToGroup(handlers.NewCallback(callbackquery.Prefix("cmd_"), CbCommand), callbackHandlerGroup)
}

// CommandHandler any unhandled commands
func CommandHandler(bot *gotgbot.Bot, ctx *ext.Context) error {
	update := ctx.EffectiveMessage
	user := ctx.EffectiveUser

	cmd := strings.ToUpper(strings.Split(strings.ToLower(strings.Fields(update.GetText())[0]), "@")[0][1:])

	text, buttons := config.GetCommand(cmd)

	text = format.BasicFormat(text, user)

	_, err := bot.SendMessage(update.Chat.Id, text, &gotgbot.SendMessageOpts{ParseMode: gotgbot.ParseModeHTML, LinkPreviewOptions: &gotgbot.LinkPreviewOptions{IsDisabled: true}, ReplyMarkup: gotgbot.InlineKeyboardMarkup{InlineKeyboard: buttons}})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func allCommand(msg *gotgbot.Message) bool {
	ents := msg.GetEntities()
	if len(ents) != 0 && ents[0].Offset == 0 && ents[0].Type != "bot_command" {
		return false
	}

	text := msg.GetText()

	if r, _ := utf8.DecodeRuneInString(text); r != '/' {
		return false
	}

	split := strings.Split(strings.ToLower(strings.Fields(text)[0]), "@")
	cmd := split[0][1:]

	return cmd != ""
}
