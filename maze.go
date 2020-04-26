package main

import (
	"bufio"
	"os"
)

type Maze struct {
	Layout  []string
	Player  *Player
	Ghosts  []*Ghost
	NumDots int
}

func NewMaze(filename string) (*Maze, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var layout []string

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		layout = append(layout, line)
	}

	return &Maze{Layout: layout}, nil
}

func (m *Maze) Populate() {
	for row, line := range m.Layout {
		for col, char := range line {
			switch char {
			case 'P':
				m.Player = NewPlayer(row, col, 3)
			case 'G':
				m.Ghosts = append(m.Ghosts, NewGhost(row, col))
			case '.':
				m.NumDots++
			}
		}
	}
}

func (m *Maze) MovePlayer(direction string) {
	m.Player.Move(direction)

	removeDot := func(row, col int) {
		m.Layout[row] = m.Layout[row][0:col] + " " + m.Layout[row][col+1:]
	}

	switch m.Layout[m.Player.Row][m.Player.Col] {
	case '.':
		m.NumDots--
		m.Player.Score++
		removeDot(m.Player.Row, m.Player.Col)
	case 'X':
		m.Player.Score += 10
		removeDot(m.Player.Row, m.Player.Col)
	}
}

func (m *Maze) MoveGhosts() {
	for _, ghost := range m.Ghosts {
		ghost.Move()

		if m.Player.Row == ghost.Row && m.Player.Col == ghost.Col {
			m.Player.Lives--
		}
	}
}
