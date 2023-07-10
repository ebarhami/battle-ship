package game

import (
	"github.com/google/uuid"
)

type Player struct {
	Id    uuid.UUID
	Board *Board
	Score int
}
