package main

import (
	"encoding/json"
	"flag"
	"os"
)

type Config struct {
	// Command line options
	ConfigFile  *string
	MazeFile    *string
	PlayerLives *int

	// Config file options
	UseEmoji bool   `json:"use_emoji"`
	Player   string `json:"player"`
	Ghost    string `json:"ghost"`
	Wall     string `json:"wall"`
	Dot      string `json:"dot"`
	Pill     string `json:"pill"`
	Death    string `json:"death"`
	Space    string `json:"space"`
}

func (c *Config) LoadFile(file string) error {
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

func (c *Config) HandleCommandLineOptions() {
	c.ConfigFile = flag.String("config-file", "config.json", "relative path to a custom configuration file")
	c.MazeFile = flag.String("maze-file", "maze01.txt", "relative path to a custom maze layout file")
	c.PlayerLives = flag.Int("player-lives", 3, "number of player lives")

	flag.Parse()
}
