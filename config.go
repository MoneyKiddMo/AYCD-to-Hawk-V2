package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"

	"github.com/fatih/color"
)

func readConfig() (*Config, error) {
	var config *Config

	configFile, err := os.Open("config.json")
	if err != nil {
		color.Red("[ERROR] config.readConfig Error Opening config.json: %s", err)
		return nil, err
	}
	defer configFile.Close()

	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		color.Red("[ERROR] config.readConfig Error Reading config.json. Closing Program in 3 seconds...")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}

	json.Unmarshal(configData, &config)
	return config, nil
}
