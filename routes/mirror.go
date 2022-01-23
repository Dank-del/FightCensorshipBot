package routes

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func mirrorHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	go func() {
		err := sendMirror(ctx.EffectiveMessage, ctx.EffectiveChat, b)
		if err != nil {
			log.Println(err)
		}
	}()
	return ext.ContinueGroups
}

func sendMirror(msg *gotgbot.Message, chat *gotgbot.Chat, bot *gotgbot.Bot) error {
	opts := &gotgbot.CopyMessageOpts{}
	if msg.ReplyMarkup != nil {
		opts.ReplyMarkup = msg.ReplyMarkup
	}
	opts.AllowSendingWithoutReply = true
	opts.CaptionEntities = msg.CaptionEntities
	opts.ReplyToMessageId = msg.MessageId
	opts.Caption = msg.Caption
	_, err := bot.CopyMessage(chat.Id, msg.Chat.Id, msg.MessageId, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func unCensorCmd(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		_, err := ctx.EffectiveMessage.Reply(b, "Reply to a message retard", nil)
		if err != nil {
			return err
		}
		return ext.EndGroups
	}
	go func() {
		err := sendMirror(ctx.EffectiveMessage.ReplyToMessage, ctx.EffectiveChat, b)
		if err != nil {
			log.Println(err)
		}
	}()
	return ext.EndGroups
}
