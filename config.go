package main

import (
	"encoding/json"
	"flag"
	"os"
	"time"
)

type Config struct {
	// Command line options
	ConfigFile       *string
	MazeFile         *string
	PlayerLives      *int
	PillDuration     *time.Duration
	PillScore        *int
	GhostDefeatScore *int
	Framerate        *time.Duration

	// Config file options
	UseEmoji  bool   `json:"use_emoji"`
	Player    string `json:"player"`
	Ghost     string `json:"ghost"`
	GhostBlue string `json:"ghost_blue"`
	Wall      string `json:"wall"`
	Dot       string `json:"dot"`
	Pill      string `json:"pill"`
	Death     string `json:"death"`
	Space     string `json:"space"`
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
	c.ConfigFile = flag.String("config-file", "./lib/config.json", "relative path to a custom configuration file")
	c.MazeFile = flag.String("maze-file", "./lib/maze01.txt", "relative path to a custom maze layout file")
	c.PlayerLives = flag.Int("player-lives", 3, "number of player lives")
	c.PillScore = flag.Int("pill-score", 10, "points scored when swallowing a pill")
	c.PillDuration = flag.Duration("pill-duration", 20*time.Second, "time for which a pill should take effect")
	c.GhostDefeatScore = flag.Int("ghost-defeat-score", 15, "points scored when defeating a ghost")
	c.Framerate = flag.Duration("framerate", 200*time.Millisecond, "speed at which to render the game")

	flag.Parse()
}
