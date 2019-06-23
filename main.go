package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var maze []string

func main() {
	// initialize the game

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

		// process movement

		// process collisions

		// check game over

		// break infinite loop
		break

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

	return nil
}

func printScreen() {
	for _, line := range maze {
		fmt.Println(line)
	}
}
