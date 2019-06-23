package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var maze []string
var player Player

func init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode in terminal: %v\n", err)
	}
}

func main() {
	// initialize the game
	defer cleanup()

	// load resources
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

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

		// process movement

		// process collisions

		// check game over

		// break infinite loop
		if input == "ESC" {
			break
		}

		// repeat
	}
}

func loadMaze() error {
	f, err := os.Open("maze01.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = Player{row, col}
			}
		}
	}

	return nil
}

func printScreen() {
	clearScreen()

	for _, line := range maze {
		fmt.Println(line)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode in terminal: %v\n", err)
	}
}
