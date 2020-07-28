package main

import (
	"fmt"
	"github.com/ValorantWLNotificationBot/util"
	"github.com/bwmarrin/discordgo"
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
