package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	HypixelToken string
	BotToken     string
)

func Run() {
	discordBot, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	//adding event handler
	discordBot.AddHandler(newMessage)

	//open session
	discordBot.Open()
	defer discordBot.Close()

	//run until closed
	fmt.Println("Bot is running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messaage
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case strings.Contains(message.Content, "bot"):
		discord.ChannelMessageSend(message.ChannelID, "Hi there!")
	case strings.Contains(message.Content, "!bw"):
		currentWeather := getBedWarInfo(message.Content)
		discord.ChannelMessageSendComplex(message.ChannelID, currentWeather)
	}

}
