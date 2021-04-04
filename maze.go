package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type Maze struct {
	Layout        []string
	Player        *Player
	Ghosts        []*Ghost
	GhostStatusMx sync.RWMutex
	NumDots       int
	PillTimer     *time.Timer
	PillTimerEnd  time.Time
	pillMx        sync.Mutex
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

	maze := &Maze{Layout: layout}

	return maze.Populate(), nil
}

func (m *Maze) Populate() *Maze {
	for row, line := range m.Layout {
		for col, char := range line {
			switch char {
			case 'P':
				m.Player = NewPlayer(row, col, cfg.PlayerLives)
			case 'G':
				m.Ghosts = append(m.Ghosts, NewGhost(row, col))
			case '.':
				m.NumDots++
			}
		}
	}

	return m
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
		m.Player.Score += cfg.PillScore
		removeDot(m.Player.Row, m.Player.Col)

		go m.processPill()
	}
}

func (m *Maze) MoveGhosts() {
	for _, ghost := range m.Ghosts {
		ghost.Move()

		if m.Player.Row == ghost.Row && m.Player.Col == ghost.Col {
			m.GhostStatusMx.Lock()
			if ghost.IsThreat {
				m.Player.LoseLife()
			} else {
				m.Player.Score += cfg.GhostDefeatScore
				ghost.Defeat()
			}
			m.GhostStatusMx.Unlock()
		}
	}
}

func (m *Maze) swapGhosts(toStatus string) {
	m.GhostStatusMx.Lock()
	defer m.GhostStatusMx.Unlock()

	for _, ghost := range m.Ghosts {
		switch toStatus {
		case "threat":
			ghost.IsThreat = true
		default:
			ghost.IsThreat = false
		}
	}
}

func (m *Maze) processPill() {
	m.pillMx.Lock()
	defer m.pillMx.Unlock()

	m.swapGhosts("safe")
	m.setPillTimer()

	m.pillMx.Unlock()
	<-m.PillTimer.C
	m.pillMx.Lock()

	m.PillTimer.Stop()
	m.swapGhosts("threat")
}

func (m *Maze) setPillTimer() {
	pillTime := cfg.PillDuration

	if m.PillTimer != nil {
		m.PillTimer.Stop()

		pillTime += time.Until(maze.PillTimerEnd)
		maze.PillTimerEnd = time.Now()
	}

	m.PillTimer = time.NewTimer(pillTime)
	m.PillTimerEnd = time.Now().Add(pillTime)
}
