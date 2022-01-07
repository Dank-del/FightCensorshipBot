package routes

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Load(d *ext.Dispatcher) {
	d.AddHandler(handlers.NewCommand("start", startHandler))
	d.AddHandler(handlers.NewMessage(func(msg *gotgbot.Message) bool {
		return msg != nil
	}, mirrorHandler))
}
