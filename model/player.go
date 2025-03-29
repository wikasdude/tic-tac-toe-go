package model

import "fmt"

type Player struct {
	Name   string
	Symbol string
}

func (p *Player) MakeMove(board *Board, row, col int) bool {
	if board.grid[row][col] == " " {
		board.grid[row][col] = p.Symbol
		return true
	}
	fmt.Println("Invalid move, try again!")
	return false
}
