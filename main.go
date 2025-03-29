package main

import (
	"fmt"
	"tic_tac_toe/game"
)

func main() {
	fmt.Println("game started")
	// player1 := &model.Player{Name: "Alice", Symbol: "X"}
	// player2 := &model.Player{Name: "Bob", Symbol: "O"}
	game := game.NewGame("human", "Alice", "X", "ai", "Bot", "O", 3)
	game.StartGame()

}
