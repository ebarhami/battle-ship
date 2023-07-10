package game

import (
	"errors"
	"fmt"

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

func (p *Player) Hit(otherPlayer *Player, coord Coordinate) error {
	otherPlayerGridSize := len(otherPlayer.Board.Grid)
	if p.Id == otherPlayer.Id {
		return errors.New("Friendly Fire")
	}
	if !coord.IsValid() || coord.X >= otherPlayerGridSize || coord.Y >= otherPlayerGridSize {
		return errors.New(fmt.Sprintf("Coordinate is invalid, Player %v has board with size %d",
			otherPlayer.Id, otherPlayerGridSize))
	}

	if otherPlayer.Board.FireCoordinate(coord) { // is hit
		p.Score++
	}
	return nil
}
