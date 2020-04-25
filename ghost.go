package main

import "math/rand"

type Ghost struct {
	Row int
	Col int
}

func NewGhost(row, col int) *Ghost {
	return &Ghost{row, col}
}

func (g *Ghost) Move(maze []string) {
	dir := drawDirection()
	g.Row, g.Col = makeMove(maze, g.Row, g.Col, dir)
}

func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "LEFT",
		3: "RIGHT",
	}

	return move[dir]
}
