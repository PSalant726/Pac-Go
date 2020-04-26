package main

import (
	"fmt"
	"time"
)

type Player struct {
	Row      int
	Col      int
	StartRow int
	StartCol int
	Lives    int
	Score    int
}

func NewPlayer(row, col, lives int) *Player {
	return &Player{row, col, row, col, lives, 0}
}

func (p *Player) Move(dir string) {
	p.Row, p.Col = makeMove(p.Row, p.Col, dir)
}

func (p *Player) LoseLife() {
	p.Lives--
	moveCursor(p.Row, p.Col)
	fmt.Print(cfg.Death)

	if p.Lives > 0 {
		time.Sleep(time.Second)

		p.Row, p.Col = p.StartRow, p.StartCol
	}
}
