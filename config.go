package main

import (
	"encoding/json"
	"os"
	"sync"
)

const configFile string = "config.json"

type Config struct {
	WssURL  string `json:"wssURL"`
	BaseURL string `json:"baseURL"`
	Account struct {
		Secret     string `json:"secret"`
		Key        string `json:"key"`
		Passphrase string `json:"passphrase"`
	} `json:"account"`
	Init       struct {
		Crypto   string `json:"crypto"`
		Currency string `json:"currency"`
	} `json:"init"`
	ConsoleLog     string `json:"consoleLog"`
}

var instance *Config
var onceConfig sync.Once

func GetConfigInstance() *Config {
	onceConfig.Do(func() {
		instance = &Config{}
		instance.loadConfiguration(configFile)
	})
	return instance
}

func (config *Config) loadConfiguration(file string) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		GetLoggerInstance().Error("In loadConfiguration: %s", err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}
