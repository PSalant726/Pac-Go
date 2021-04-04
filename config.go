package main

import (
	"encoding/json"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

type Config struct {
	// Command line options
	ConfigFile       string
	MazeFile         string
	PlayerLives      int
	PillDuration     time.Duration
	PillScore        int
	GhostDefeatScore int
	Framerate        time.Duration

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
	flag.StringVarP(&c.ConfigFile, "config-file", "c", "./lib/config.json", "relative path to a custom configuration file")
	flag.StringVarP(&c.MazeFile, "maze-file", "m", "./lib/maze01.txt", "relative path to a custom maze layout file")
	flag.IntVarP(&c.PlayerLives, "player-lives", "l", 3, "number of player lives")
	flag.IntVarP(&c.PillScore, "pill-score", "p", 10, "points scored when swallowing a pill")
	flag.DurationVarP(&c.PillDuration, "pill-duration", "d", 20*time.Second, "time for which a pill should take effect")
	flag.IntVarP(&c.GhostDefeatScore, "ghost-defeat-score", "g", 15, "points scored when defeating a ghost")
	flag.DurationVarP(&c.Framerate, "framerate", "f", 200*time.Millisecond, "speed at which to render the game")

	flag.Parse()
}
