package main

import (
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
	cfg.HandleCommandLineOptions()

	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	if err := cbTerm.Run(); err != nil {
		log.Fatalf("Unable to activate cbreak mode in terminal: %v\n", err)
	}
}

func main() {
	var err error

	// load resources
	if maze, err = NewMaze(cfg.MazeFile); err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	if err = cfg.LoadFile(cfg.ConfigFile); err != nil {
		log.Println("failed to load configuration:", err)
		return
	}

	playGame()
	cleanup()
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	if err := cookedTerm.Run(); err != nil {
		log.Fatalf("Unable to activate cooked mode in terminal: %v\n", err)
	}
}

func playGame() {
	// process input (async)
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := readInput()
			if err != nil {
				log.Println("failed to read user input:", err)
				ch <- "ESC"
			}

			ch <- input
		}
	}(input)

	for {
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

		if msg, ok := isGameOver(); ok {
			printScreen()
			updatePlayerMessage(msg)

			break
		}

		time.Sleep(cfg.Framerate)
	}
}
