package main

import (
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
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode in terminal: %v\n", err)
	}
}

func main() {
	var err error

	// initialize the game
	defer cleanup()

	// load resources
	maze, err = NewMaze(Blueprint{"maze01.txt"})
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	maze.Populate()

	if err = cfg.Load("config.json"); err != nil {
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

			moveCursor(len(maze.Layout)+2, 0)
			fmt.Println("\n\t  Game Over")

			break
		} else if maze.NumDots == 0 {
			fmt.Println("\nCongratulations! You win!")
			break
		}

		// repeat
		time.Sleep(200 * time.Millisecond)
	}
}
