package core

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"os"
)

type Secret struct {
	discordClientSecret string
}

func getSecrets() Secret {
	var secret Secret
	jsonFile, err := os.Open("../.secrets/secrets.json")
	if err != nil {
		fmt.Println("Error reading in secrets.json:", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &secret)
	return secret
}

func main() {
	secret := getSecrets()
	if secret.discordClientSecret == "" {
		fmt.Println("Unable to read in discord client secret.")
		return
	}
	discord, err := discordgo.New("Bot " + secret.discordClientSecret)
	if err != nil {
		fmt.Println("Error creating Discord sesssion:", err)
	}

	discord.Close()
}