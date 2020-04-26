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

type Blueprint struct {
	Filename string
}

func NewMaze(blueprint Blueprint) (*Maze, error) {
	f, err := os.Open(blueprint.Filename)
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
	m.Player.Move(direction, m.Layout)

	switch m.Layout[m.Player.Row][m.Player.Col] {
	case '.':
		m.NumDots--
		m.Player.Score++
		// Remove dot from the maze
		m.Layout[m.Player.Row] = m.Layout[m.Player.Row][0:m.Player.Col] + " " + m.Layout[m.Player.Row][m.Player.Col+1:]
	}
}

func (m *Maze) MoveGhosts() {
	for _, ghost := range m.Ghosts {
		ghost.Move(m.Layout)

		if m.Player.Row == ghost.Row && m.Player.Col == ghost.Col {
			m.Player.Lives--
		}
	}
}
