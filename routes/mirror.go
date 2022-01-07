package routes

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"log"
)

func mirrorHandler(b *gotgbot.Bot, ctx *ext.Context) error {
	chat := ctx.EffectiveChat
	msg := ctx.EffectiveMessage
	opts := &gotgbot.CopyMessageOpts{}
	if msg.ReplyMarkup != nil {
		opts.ReplyMarkup = msg.ReplyMarkup
	}
	opts.AllowSendingWithoutReply = true
	opts.CaptionEntities = msg.CaptionEntities
	opts.ReplyToMessageId = msg.MessageId
	opts.Caption = msg.Caption
	_, err := b.CopyMessage(chat.Id, msg.Chat.Id, msg.MessageId, opts)
	if err != nil {
		log.Println(err)
		return err
	}
	return ext.ContinueGroups
}
