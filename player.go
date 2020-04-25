package main

type Player struct {
	Row   int
	Col   int
	Lives int
}

func NewPlayer(row, col, lives int) *Player {
	return &Player{row, col, lives}
}

func (p *Player) Move(dir string) {
	p.Row, p.Col = makeMove(p.Row, p.Col, dir)

	switch maze[p.Row][p.Col] {
	case '.':
		numDots--
		score++
		// Remove dot from the maze
		maze[p.Row] = maze[p.Row][0:p.Col] + " " + maze[p.Row][p.Col+1:]
	}
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow--

		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow++

		if newRow == len(maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol++

		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol--

		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return newRow, newCol
}
