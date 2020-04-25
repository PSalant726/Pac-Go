package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var (
	maze    []string
	player  *Player
	ghosts  []*Ghost
	score   int
	numDots int
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
		player.Move(input)

		for _, ghost := range ghosts {
			ghost.Move()
		}

		// process collisions
		for _, ghost := range ghosts {
			if player.Row == ghost.Row && player.Col == ghost.Col {
				player.Lives--
			}
		}

		// check game over
		if input == "ESC" || player.Lives <= 0 {
			fmt.Println("\n\t  Game Over")
			break
		} else if numDots == 0 {
			fmt.Println("\nCongratulations! You win!")
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
				player = NewPlayer(row, col, 3)
			case 'G':
				ghosts = append(ghosts, NewGhost(row, col))
			case '.':
				numDots++
			}
		}
	}

	return nil
}

func printScreen() {
	clearScreen()

	for _, line := range maze {
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

	moveCursor(player.Row, player.Col)
	fmt.Printf("P")

	for _, g := range ghosts {
		simpleansi.MoveCursor(g.Row, g.Col)
		fmt.Print("G")
	}

	simpleansi.MoveCursor(len(maze)+1, 0)
	fmt.Println("  Score:", score, "\t  Lives:", player.Lives)
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
