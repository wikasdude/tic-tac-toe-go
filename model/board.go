package model

import (
	"fmt"
	"sync"
)

var (
	board *Board
	once  sync.Once
)

type Board struct {
	grid [][]string
	size int
}

func NewBoard(size int) *Board {
	//var board Board
	once.Do(func() {
		board = &Board{size: size,
			grid: make([][]string, size),
		}
		for val := range board.grid {
			board.grid[val] = make([]string, size)
			for j := range board.grid[val] {
				board.grid[val][j] = " "
			}

		}

	})

	return board
}
func (b *Board) PrintBoard() {
	for key, val := range b.grid {
		fmt.Println(key, val)

	}
}

func (b *Board) IsFull() bool {

	for _, val1 := range b.grid {
		for _, val2 := range val1 {
			if val2 == " " {
				return false
			}

		}
	}

	return true
}
func (b *Board) CheckWin(symbol string) bool {
	for i := 0; i < b.size; i++ {
		if checkLine(b.grid[i], symbol) || checkLine(getColumn(b.grid, i), symbol) {
			return true
		}
	}
	if checkLine(getDiagonal(b.grid, true), symbol) || checkLine(getDiagonal(b.grid, false), symbol) {
		return true
	}
	return false
}
func checkLine(line []string, symbol string) bool {
	for _, cell := range line {
		if cell != symbol {
			return false
		}
	}
	return true
}
func getColumn(grid [][]string, col int) []string {
	column := []string{}
	for _, row := range grid {
		column = append(column, row[col])
	}
	return column
}
func getDiagonal(grid [][]string, mainDiagonal bool) []string {
	diagonal := []string{}
	for i := 0; i < len(grid); i++ {
		if mainDiagonal {
			diagonal = append(diagonal, grid[i][i])
		} else {
			diagonal = append(diagonal, grid[i][len(grid)-1-i])
		}
	}
	return diagonal
}
