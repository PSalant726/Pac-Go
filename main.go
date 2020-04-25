package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var (
	maze  *Maze
	score int
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

	// game loop
	for {
		// update screen
		printScreen()

		// process input
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v\n", err)
			break
		}

		// process movement & collisions
		maze.MovePlayer(input)
		maze.MoveGhosts()

		// check game over
		if input == "ESC" || maze.Player.Lives <= 0 {
			fmt.Println("\n\t  Game Over")
			break
		} else if maze.NumDots == 0 {
			fmt.Println("\nCongratulations! You win!")
			break
		}

		// repeat
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
	fmt.Println("  Score:", score, "\t  Lives:", maze.Player.Lives)
}
