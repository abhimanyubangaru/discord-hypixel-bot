package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const URL string = "https://api.hypixel.net/player?"

type HypixelData struct {
	Player struct {
		DisplayName string `json:"displayname"`
		Stats       struct {
			BedWars struct {
				Kills       int `json:"kills_bedwars"`
				GamesPlayed int `json:"games_played_bedwars"`
				Wins        int `json:"wins_bedwars"`
				Deaths      int `json:"void_deaths_bedwars"`
			} `json:"BedWars"`
		} `json:"stats"`
	} `json:"player"`
}

func getBedWarInfo(message string) *discordgo.MessageSend {
	arr := strings.Fields(message)
	playerName := arr[len(arr)-1]
	if playerName == "" {
		return &discordgo.MessageSend{
			Content: "Sorry that ZIP code does not look right",
		}
	}

	hypixelURL := fmt.Sprintf("%skey=%s&name=%s", URL, HypixelToken, playerName)
	fmt.Println("Hypixel url: " + hypixelURL)

	//create new http client and set timeout
	client := http.Client{Timeout: 5 * time.Second}

	//Query Hypixel
	response, err := client.Get(hypixelURL)
	if err != nil {
		return &discordgo.MessageSend{
			Content: "Sorry, there was an error trying to get the weather",
		}
	}

	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var data HypixelData
	json.Unmarshal(body, &data) // peep the & as seen in unmarshalled in config.go

	kills := strconv.Itoa(data.Player.Stats.BedWars.Kills)
	deaths := strconv.Itoa(data.Player.Stats.BedWars.Deaths)
	gamesPlayed := strconv.Itoa(data.Player.Stats.BedWars.GamesPlayed)
	playerWins := strconv.Itoa(data.Player.Stats.BedWars.Wins)
	name := data.Player.DisplayName

	killDeathRatio := strconv.FormatFloat(float64(data.Player.Stats.BedWars.Kills)/float64(data.Player.Stats.BedWars.Deaths), 'f', -1, 32)
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{{
			Type:        discordgo.EmbedTypeRich,
			Title:       "Bedwars Stats for " + name,
			Description: "K/D: " + killDeathRatio,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Kills",
					Value:  kills,
					Inline: true,
				},
				{
					Name:   "Deaths",
					Value:  deaths,
					Inline: true,
				},
				{
					Name:   "Games Played",
					Value:  gamesPlayed + " Games Played",
					Inline: true,
				},
				{
					Name:   "Wins",
					Value:  playerWins + " wins",
					Inline: true,
				},
			},
		},
		},
	}
	return embed
}
