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

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode in terminal: %v\n", err)
	}
}

func main() {
	defer cleanup()

	var err error

	// load resources
	maze, err = NewMaze(*cfg.MazeFile)
	if err != nil {
		log.Println("failed to load maze:", err)
		return
	}

	if err = cfg.LoadFile(*cfg.ConfigFile); err != nil {
		log.Println("failed to load configuration:", err)
		return
	}

	playGame()
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
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

		if isGameOver() {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}
}
