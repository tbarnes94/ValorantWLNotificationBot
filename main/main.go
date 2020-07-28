package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ValorantWLNotificationBot/util"
)



func main() {
	secret := util.GetSecrets()
	if secret.DiscordClientSecret == "" {
		fmt.Println("Unable to read in discord client secret.")
		return
	}
	discord, err := discordgo.New("Bot " + secret.DiscordClientSecret)
	if err != nil {
		fmt.Println("Error creating Discord sesssion:", err)
	}

	discord.Close()
}