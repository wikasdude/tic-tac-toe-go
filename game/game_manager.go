package game

import (
	"fmt"
	"tic_tac_toe/model"
)

type Game struct {
	board         *model.Board
	players       []*model.Player
	currentPlayer int
}

func NewGame(player1, player2 *model.Player, boardSize int) *Game {
	return &Game{
		board:   model.NewBoard(boardSize),
		players: []*model.Player{player1, player2},
	}
}
func (g *Game) StartGame() {
	fmt.Println("Game started!")
	g.board.PrintBoard()

	for {
		player := g.players[g.currentPlayer]
		var row, col int
		fmt.Printf("%s, enter your move (row col): ", player.Name)
		fmt.Scan(&row, &col)
		if player.MakeMove(g.board, row, col) {
			g.board.PrintBoard()
			if g.board.CheckWin(player.Symbol) {
				fmt.Printf("%s wins!\n", player.Name)
				return
			}
			if g.board.IsFull() {
				fmt.Println("It's a draw!")
				return
			}
			g.currentPlayer = (g.currentPlayer + 1) % 2
		}
	}
}
