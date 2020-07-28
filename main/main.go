package main

import (
	"fmt"
	"github.com/ValorantWLNotificationBot/util"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	uuid "github.com/satori/go.uuid"
)

func main() {
	secret := util.GetSecrets()
	if secret.DiscordClientToken == "" {
		fmt.Println("Unable to read in discord client secret.")
		return
	}
	discord, err := discordgo.New("Bot " + secret.DiscordClientToken)
	if err != nil {
		fmt.Println("Error creating Discord sesssion:", err)
	}

	player := "Cerebrus"
	region := "na1"
	url := util.GetURLForUUID(player, region)

	res, err := discord.Client.Get(url)
	if err != nil {
		fmt.Println("Error on GET to " + url + ":", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error on reading HTTP Response Body:", err)
	}
	id, err := uuid.FromString(string(body[7:43]))

	fmt.Println(id)
	st, err := discord.Channel(util.ValorantChannelID)
	if err != nil {
		fmt.Println("Error on connecting to Valorant channel:", err)
	}
	fmt.Println(st.Name)

	content := fmt.Sprintf(util.Message, player, id.String())

	out, err := discord.Gateway()
	if err != nil {
		fmt.Println("Error on connecting bot to gateway:", err)
	}
	fmt.Println(out)

	message, err := discord.ChannelMessageSend(util.ValorantChannelID, content)
	if err != nil {
		fmt.Println("Error on sending message to Valorant channel:", err)
	}
	fmt.Println(message)

	discord.Close()
}
