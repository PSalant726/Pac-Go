package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	ConfigFile  *string
	MazeFile    *string
	PlayerLives *int

	UseEmoji bool   `json:"use_emoji"`
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
}

func (c *Config) Load(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	if err = decoder.Decode(&c); err != nil {
		return err
	}

	return nil
}
