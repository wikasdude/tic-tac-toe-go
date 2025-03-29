package factory

import "tic_tac_toe/model"

const (
	Human = "human"
	AI    = "ai"
)

func NewPlayerFactory(playerType, name, symbol string) *model.Player {
	switch playerType {
	case Human:
		return &model.Player{Name: name, Symbol: symbol}
	case AI:
		return &model.Player{Name: "AI-" + name, Symbol: symbol} // AI Player (Basic)
	default:
		return nil
	}
}
