package game

import (
	"github.com/google/uuid"
)

type Player struct {
	Id    uuid.UUID
	Board *Board
	Score int
}

func NewPlayer(boardSize int) *Player {
	board := NewBoard(boardSize)
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil
	}
	return &Player{
		Id:    uuid,
		Board: board,
	}
}
