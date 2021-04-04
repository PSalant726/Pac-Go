package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/danicat/simpleansi"
)

func getLivesAsEmoji() string {
	buf := bytes.Buffer{}

	for i := maze.Player.Lives; i > 0; i-- {
		buf.WriteString(cfg.Player + " ")
	}

	return buf.String()
}

func isGameOver() (string, bool) {
	if maze.Player.Lives <= 0 {
		moveCursor(maze.Player.Row, maze.Player.Col)
		fmt.Print(cfg.Death)

		return "Game Over", true
	} else if maze.NumDots == 0 {
		return "Congratulations! You win!", true
	}

	return "", false
}

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

		maze.GhostStatusMx.RLock()
		if ghost.IsThreat {
			fmt.Print(cfg.Ghost)
		} else {
			fmt.Print(cfg.GhostBlue)
		}
		maze.GhostStatusMx.RUnlock()
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
	livesRemaining := strconv.Itoa(maze.Player.Lives)
	if cfg.UseEmoji {
		livesRemaining = getLivesAsEmoji()
	}

	var pillTimeRemaining string
	if pillTime := time.Until(maze.PillTimerEnd); pillTime > 0 {
		pillTimeRemaining = fmt.Sprintf("\nPill Time Remaining: %v", pillTime.Round(time.Second))
	}

	moveCursor(len(maze.Layout)+1, 0)
	fmt.Println("Score:", maze.Player.Score, "\nLives:", livesRemaining, pillTimeRemaining)

	moveCursor(len(maze.Layout)+4, 0)
	fmt.Println(message)
}
