package main

import (
	"FightCensorshipBot/core"
	"FightCensorshipBot/routes"
	"log"
)

func main() {
	core.Conf = core.NewBotConfig()
	core.Conf.ReadFile("config.json")
	_, up, err := core.BotInit(core.Conf)
	if err != nil {
		log.Println(err)
	}
	routes.Load(up.Dispatcher)
	up.Idle()
}
