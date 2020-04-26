package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/danicat/simpleansi"
)

var maze *Maze

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

func printScreen() {
	clearScreen()

	for _, line := range maze.Layout {
		for _, char := range line {
			switch char {
			case '#':
				fallthrough
			case '.':
				fmt.Printf("%c", char)
			default:
				fmt.Printf(" ")
			}
		}

		fmt.Print("\n")
	}

	moveCursor(maze.Player.Row, maze.Player.Col)
	fmt.Printf("P")

	for _, ghost := range maze.Ghosts {
		simpleansi.MoveCursor(ghost.Row, ghost.Col)
		fmt.Print("G")
	}

	simpleansi.MoveCursor(len(maze.Layout)+1, 0)
	fmt.Println("  Score:", maze.Player.Score, "\t  Lives:", maze.Player.Lives)
}
