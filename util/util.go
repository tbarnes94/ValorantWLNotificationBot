package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Secret struct {
	DiscordClientSecret string
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
