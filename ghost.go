package main

import "math/rand"

type Ghost struct {
	Row      int
	Col      int
	StartRow int
	StartCol int
	IsThreat bool
}

func NewGhost(row, col int) *Ghost {
	return &Ghost{row, col, row, col, true}
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

func (g *Ghost) Move() {
	dir := drawDirection()
	g.Row, g.Col = makeMove(g.Row, g.Col, dir)
}

func (g *Ghost) Defeat() {
	g.Row, g.Col = g.StartRow, g.StartCol
	g.IsThreat = true
}
