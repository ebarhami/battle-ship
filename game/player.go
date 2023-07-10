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
	if p.Id == otherPlayer.Id {
		return errors.New("Friendly Fire")
	}
	if !coord.IsValid() || coord.X >= len(otherPlayer.Board.Grid) || coord.Y >= len(otherPlayer.Board.Grid) {
		return errors.New(fmt.Sprintf("Coordinate is invalid, Player %v has board with size %d",
			otherPlayer.Id, len(otherPlayer.Board.Grid)))
	}

	if otherPlayer.Board.FireCoordinate(coord) { // is hit
		p.Score++
	}
	return nil
}
