package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	maze *Maze
	cfg  Config
)

func init() {
	cfg.ConfigFile = flag.String("config-file", "config.json", "relative path to a custom configuration file")
	cfg.MazeFile = flag.String("maze-file", "maze01.txt", "relative path to a custom maze layout file")
	cfg.PlayerLives = flag.Int("player-lives", 3, "number of player lives")
	flag.Parse()

	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode in terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode in terminal: %v\n", err)
	}
}

func main() {
	defer cleanup()

	var err error

	// load resources
	maze, err = NewMaze(*cfg.MazeFile)
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	if err = cfg.Load(*cfg.ConfigFile); err != nil {
		log.Println("failed to load configuration:", err)
		return
	}

	// process input (async)
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := readInput()
			if err != nil {
				log.Println("Error reading input:", err)
				ch <- "ESC"
			}

			ch <- input
		}
	}(input)

	// game loop
	for {
		// update screen
		printScreen()

		// process movement & collisions
		select {
		case inp := <-input:
			if inp == "ESC" {
				maze.Player.Lives = 0
			}

			maze.MovePlayer(inp)
		default:
		}

		maze.MoveGhosts()

		// check game over
		if maze.Player.Lives <= 0 {
			moveCursor(maze.Player.Row, maze.Player.Col)
			fmt.Print(cfg.Death)

			updatePlayerMessage("Game Over")

			break
		} else if maze.NumDots == 0 {
			updatePlayerMessage("Congratulations! You win!")
			break
		}

		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
