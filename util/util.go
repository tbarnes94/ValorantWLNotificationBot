package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const URL_ValorantGetUUID = "https://valorant.iesdev.com/player/"
const GeneralChannelID = "258438663438991360"
const ValorantChannelID = "699325128080883804"
const Message = "Hello, %s, your Valorant UUID is %s."

type Secret struct {
	DiscordClientToken string
}

func GetURLForUUID(player, region string) string {
	return URL_ValorantGetUUID + strings.ToLower(player) + "-" + region
}

func GetSecrets() Secret {
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
