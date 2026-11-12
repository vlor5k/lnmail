package config

import (
	"encoding/json/v2"
	"fmt"
	"os"
)

type ConfigType struct {
	// Domains                []string

	// Random emails
	RandomEmailsEnabled   bool
	RandomEmailsPriceSats int64

	// Custom emails
	CustomEmailsEnabled   bool
	CustomEmailsPriceSats int64

	// Ln
	LightningConnectionUri string
	LightningNode          string
	LightningRune          string

	// Tor
	OnionsiteUrl string

	// Landing config
	ShowStatsOnLandingPage bool
}

var Config ConfigType

func LoadConfig() {
	data, err := os.ReadFile("config.json")

	if err != nil {
		panic("Couldn't read config.json")
	}

	err = json.Unmarshal(data, &Config)

	if err != nil {
		panic("Couldn't decode config.json")
	}

	if Config.CustomEmailsEnabled && Config.CustomEmailsPriceSats == 0 {
		fmt.Println("Warning: Custom emails are ENABLED but have a price of 0")
	}

	if Config.RandomEmailsEnabled && Config.RandomEmailsPriceSats == 0 {
		fmt.Println("Warning: Random emails are ENABLED but have a price of 0")
	}

	fmt.Printf("%+v\n", Config)
}
