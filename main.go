package main

import (
	"fmt"
	"tic_tac_toe/game"
	"tic_tac_toe/model"
)

func main() {
	fmt.Println("game started")
	player1 := &model.Player{Name: "Alice", Symbol: "X"}
	player2 := &model.Player{Name: "Bob", Symbol: "O"}
	game := game.NewGame(player1, player2, 3)
	game.StartGame()

}
