package config

import (
	"encoding/json"
	"os"
)

type Shortcuts struct {
	Menu            string `json:"menu"`
	GotoDashboard   string `json:"goto_dashboard"`
	GotoLogin       string `json:"goto_login"`
	Refresh         string `json:"refresh"`
	CloseMenu       string `json:"close_menu"`
	Quit            string `json:"quit"`
	GoToInstruments string `json:"goto_instruments"`
	GoToQuotes      string `json:"goto_quotes"`
}

type Config struct {
	Shortcuts Shortcuts `json:"shortcuts"`
}

var C Config // global loaded config

func LoadConfig() error {
	data, err := os.ReadFile("config/keys.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &C)
}
