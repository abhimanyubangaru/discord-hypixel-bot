package main

import (
	"discord-hypixel-bot/bot"
	"discord-hypixel-bot/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.BotToken = config.BotToken
	bot.HypixelToken = config.HypixelToken
	bot.Run()
}
