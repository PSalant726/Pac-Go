package main

type Player struct {
	Row   int
	Col   int
	Lives int
	Score int
}

func NewPlayer(row, col, lives int) *Player {
	return &Player{row, col, lives, 0}
}

func (p *Player) Move(dir string, maze []string) {
	p.Row, p.Col = makeMove(maze, p.Row, p.Col, dir)
}
