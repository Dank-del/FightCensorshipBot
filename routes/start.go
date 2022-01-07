package routes

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func startHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	msg := ctx.EffectiveMessage
	_, err := msg.Reply(b, "*Hi*, I'm a bot made to get away with censored posts with telegram", &gotgbot.SendMessageOpts{
		ReplyMarkup: gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{{
					Text: "Take queries here",
					Url:  "https://t.me/chiruzon",
				}},
			},
		},
		ParseMode: "markdown",
	})
	if err != nil {
		return err
	}
	return ext.EndGroups
}
