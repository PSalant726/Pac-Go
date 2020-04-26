package main

import (
	"fmt"
	"os"

	"github.com/danicat/simpleansi"
)

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow--

		if newRow < 0 {
			newRow = len(maze.Layout) - 1
		}
	case "DOWN":
		newRow++

		if newRow == len(maze.Layout)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol++

		if newCol == len(maze.Layout[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol--

		if newCol < 0 {
			newCol = len(maze.Layout[0]) - 1
		}
	}

	if maze.Layout[newRow][newCol] == '#' {
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
			case 'X':
				fmt.Printf(cfg.Pill)
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

	updatePlayerMessage("")
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

func updatePlayerMessage(message string) {
	moveCursor(len(maze.Layout)+1, 2)
	fmt.Println("Score:", maze.Player.Score)

	skewFactor := 2
	if cfg.UseEmoji {
		skewFactor = 1
	}

	moveCursor(len(maze.Layout)+1, len(maze.Layout[0])-5*skewFactor)
	fmt.Println("Lives:", maze.Player.Lives)

	moveCursor(len(maze.Layout)+3, (len(maze.Layout[0])-len(message))/2)
	fmt.Println(message)
}
