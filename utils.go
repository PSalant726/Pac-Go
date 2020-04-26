package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode in terminal: %v\n", err)
	}
}

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func makeMove(maze []string, oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow--

		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow++

		if newRow == len(maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol++

		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol--

		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return newRow, newCol
}

func moveCursor(row, col int) {
	if cfg.UseEmoji {
		simpleansi.MoveCursor(row, col*2)
	} else {
		simpleansi.MoveCursor(row, col)
	}
}

func printScreen() {
	simpleansi.ClearScreen()

	for _, line := range maze.Layout {
		for _, char := range line {
			switch char {
			case '#':
				fmt.Print(simpleansi.WithBlueBackground(cfg.Wall))
			case '.':
				fmt.Printf(cfg.Dot)
			default:
				fmt.Printf(cfg.Space)
			}
		}

		fmt.Println()
	}

	moveCursor(maze.Player.Row, maze.Player.Col)
	fmt.Printf(cfg.Player)

	for _, ghost := range maze.Ghosts {
		moveCursor(ghost.Row, ghost.Col)
		fmt.Print(cfg.Ghost)
	}

	moveCursor(len(maze.Layout)+1, 0)
	fmt.Println("  Score:", maze.Player.Score, "\t  Lives:", maze.Player.Lives)
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if buffer[0] == 0x1b {
		if cnt == 1 {
			return "ESC", nil
		} else if cnt >= 3 && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}