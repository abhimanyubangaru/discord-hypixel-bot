package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// declaring all variables
var (
	BotToken     string
	BotPrefix    string
	HypixelToken string

	config *configStruct
)

type configStruct struct {
	BotToken     string `json:"DiscordBotToken"`
	BotPrefix    string `json:"BotPrefix"`
	HypixelToken string `json:"HypixelToken"`
}

func ReadConfig() error {
	fmt.Println("Reading config file...")

	file, err := os.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(string(file))

	//unmarshalling json requires reference to config
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Bot Token: ", config.BotToken)
	fmt.Println("Bot Prefix: ", config.BotPrefix)
	fmt.Println("Hypixel Token: ", config.HypixelToken)

	//assigning all proper values
	BotToken = config.BotToken
	BotPrefix = config.BotPrefix
	HypixelToken = config.HypixelToken

	return nil
}
